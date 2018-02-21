package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/appscode/go/log"
	"github.com/appscode/grpc-go-addons/cors"
	gwrt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	Config
}

func (s *Server) Run(stopCh <-chan struct{}) error {
	if s.UseTLS() {
		go s.ServeHTTPS()
	}

	listener, err := net.Listen("tcp", s.PlaintextAddr)
	if err != nil {
		return err
	}

	m := cmux.New(listener)

	// We first match the connection against HTTP2 fields. If matched, the
	// connection will be sent through the "grpcl" listener.
	grpcl := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	// Otherwise, we match it againts HTTP1 methods. If matched,
	// it is sent through the "httpl" listener.
	httpl := m.Match(cmux.Any())

	// Then we used the muxed listeners.
	go func() {
		log.Infoln("[GRPCSERVER] Starting gRPC Server at addr", grpcl.Addr())
		log.Fatalln("[GRPCSERVER] gRPC Server failed:", s.newGRPCServer(false).Serve(grpcl))
	}()
	go func() {
		gwMux := s.NewGatewayMux(httpl, false)
		log.Infoln("[PROXYSERVER] Sarting Proxy Server at port", httpl.Addr())
		srv := &http.Server{
			Addr:         httpl.Addr().String(),
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gwMux.ServeHTTP(w, r)
			}),
		}
		log.Fatalln("[PROXYSERVER] Proxy Server failed:", srv.Serve(httpl))
	}()

	return m.Serve()
}

func (s *Server) newGRPCServer(useTLS bool) *grpc.Server {
	var gRPCServer *grpc.Server
	if useTLS {
		creds, err := credentials.NewServerTLSFromFile(s.CertFile, s.KeyFile)
		if err != nil {
			log.Fatalln(err)
		}
		s.grpcOptions = append(s.grpcOptions, grpc.Creds(creds))
	}
	gRPCServer = grpc.NewServer(s.grpcOptions...)
	s.grpcRegistry.ApplyTo(gRPCServer)
	return gRPCServer
}

func (s *Server) NewGatewayMux(l net.Listener, useTLS bool) *gwrt.ServeMux {
	gwMux := gwrt.NewServeMux(s.gwMuxOptions...)
	var grpcDialOptions []grpc.DialOption
	if useTLS {
		grpcDialOptions = []grpc.DialOption{
			grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, s.APIDomain)),
		}
	} else {
		grpcDialOptions = []grpc.DialOption{grpc.WithInsecure()}
	}
	if s.EnableCORS {
		h := cors.NewHandler(s.corsRegistry, cors.OriginHost(s.CORSOriginHost), cors.AllowSubdomain(s.CORSAllowSubdomain))
		h.RegisterHandler(gwMux)
	}

	addr := l.Addr().String()
	addr = "127.0.0.1" + addr[strings.LastIndex(addr, ":"):]
	s.proxyRegistry.ApplyTo(gwMux, addr, grpcDialOptions)
	return gwMux
}

func (s *Server) ServeHTTPS() {
	l, err := net.Listen("tcp", s.SecureAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Load certificates.
	certificate, err := tls.LoadX509KeyPair(s.CertFile, s.KeyFile)
	if err != nil {
		log.Fatalln(err)
	}
	/*
		Ref:
		 - https://blog.cloudflare.com/exposing-go-on-the-internet/
		 - http://www.bite-code.com/2015/06/25/tls-mutual-auth-in-golang/
		 - http://www.hydrogen18.com/blog/your-own-pki-tls-golang.html
	*/
	tlsConfig := &tls.Config{
		Certificates:             []tls.Certificate{certificate},
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
		SessionTicketsDisabled:   true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		ClientAuth: tls.VerifyClientCertIfGiven,
		NextProtos: []string{"h2", "http/1.1"},
	}
	if s.CACertFile != "" {
		caCert, err := ioutil.ReadFile(s.CACertFile)
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.ClientCAs = caCertPool
	}

	grpcServer := s.newGRPCServer(true)
	gwMux := s.NewGatewayMux(l, true)

	srv := &http.Server{
		Addr:         s.SecureAddr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
				grpcServer.ServeHTTP(w, r)
			} else {
				gwMux.ServeHTTP(w, r)
			}
		}),
		TLSConfig: tlsConfig,
	}

	log.Infoln("[HTTP2] Starting HTTP2 Server at port", l.Addr().String())
	log.Fatalln("[HTTP2] HTTP2 Server failed:", srv.Serve(tls.NewListener(l, tlsConfig)))
}
