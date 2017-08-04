package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	gort "runtime"
	"strings"
	"time"

	"github.com/appscode/go/runtime"
	stringz "github.com/appscode/go/strings"
	"github.com/appscode/log"
	"github.com/appscode/wheel/pkg/server/cmd/options"
	"github.com/appscode/wheel/pkg/server/endpoints"
	"github.com/appscode/wheel/pkg/server/interceptors"
	goprom "github.com/grpc-ecosystem/go-grpc-prometheus"
	gwrt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type apiServer struct {
	SecureAddr               string
	PlaintextAddr            string
	APIDomain                string
	CACertFile               string
	CertFile                 string
	KeyFile                  string
	EnableJavaClient         bool
	EnableCORS               bool
	CORSOriginHost           string
	CORSOriginAllowSubdomain bool
	GRPCEndpoints            endpoints.GRPCEndpoints
	ProxyEndpoints           endpoints.ProxyEndPoints
}

func (s *apiServer) UseTLS() bool {
	return !(s.CACertFile == "" && s.CertFile == "" && s.KeyFile == "")
}

func (s *apiServer) ListenAndServe() {
	if s.UseTLS() {
		go s.ServeHTTPS()
	}

	plaintextListener, err := net.Listen("tcp", s.PlaintextAddr)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(plaintextListener)

	// We first match the connection against HTTP2 fields. If matched, the
	// connection will be sent through the "grpcl" listener.
	var grpcl net.Listener
	if s.EnableJavaClient {
		grpcl = m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	} else {
		grpcl = m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	}

	// Otherwise, we match it againts HTTP1 methods. If matched,
	// it is sent through the "httpl" listener.
	httpl := m.Match(cmux.Any())

	// Then we used the muxed listeners.
	go func() {
		defer runtime.HandleCrash()

		log.Infoln("[GRPCSERVER] Starting gRPC Server at addr", grpcl.Addr())
		log.Fatalln("[GRPCSERVER] gRPC Server failed:", s.newGRPCServer(false).Serve(grpcl))
	}()
	go func() {
		defer runtime.HandleCrash()

		log.Infoln("[PROXYSERVER] Sarting Proxy Server at port", httpl.Addr())
		srv := &http.Server{
			Addr:         httpl.Addr().String(),
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Handler:      s.newGatewayMux(httpl, false),
		}
		log.Fatalln("[PROXYSERVER] Proxy Server failed:", srv.Serve(httpl))
	}()

	log.Fatalln(m.Serve())
}

func (s *apiServer) newGRPCServer(useTLS bool) *grpc.Server {
	var gRPCServer *grpc.Server
	if useTLS {
		creds, err := credentials.NewServerTLSFromFile(s.CertFile, s.KeyFile)
		if err != nil {
			log.Fatalln(err)
		}
		gRPCServer = grpc.NewServer(grpc.UnaryInterceptor(interceptors.NewUnaryInterceptor(s.EnableCORS, s.CORSOriginHost, s.CORSOriginAllowSubdomain)), grpc.Creds(creds))
	} else {
		gRPCServer = grpc.NewServer(grpc.UnaryInterceptor(interceptors.NewUnaryInterceptor(s.EnableCORS, s.CORSOriginHost, s.CORSOriginAllowSubdomain)))
	}

	// Register gRPC Prometheus monitoring interceptors
	goprom.Register(gRPCServer)
	// Enable Time Histogram
	goprom.EnableHandlingTimeHistogram()

	for _, endpoint := range s.GRPCEndpoints {
		log.Infoln("Registering server:", reflect.TypeOf(endpoint.Server))
		endpoints.RegisterGRPC(endpoint.RegisterFunc, gRPCServer, endpoint.Server)
	}
	return gRPCServer
}

/*
gwrt.EqualFoldMatcher("Origin"),
gwrt.EqualFoldMatcher("Cookie"),
gwrt.EqualFoldMatcher("X-Phabricator-Csrf"),
gwrt.PrefixFoldMatcher("access-control-"),
gwrt.EqualFoldMatcher("vary"),
gwrt.EqualFoldMatcher("x-content-type-options"),
gwrt.PrefixFoldMatcher("x-ratelimit-"),
*/
func (s *apiServer) newGatewayMux(l net.Listener, useTLS bool) *gwrt.ServeMux {
	gwMux := gwrt.NewServeMux(
		gwrt.WithIncomingHeaderMatcher(func(h string) (string, bool) {
			if stringz.PrefixFold(h, "access-control-request-") ||
				strings.EqualFold(h, "Origin") ||
				strings.EqualFold(h, "Cookie") ||
				strings.EqualFold(h, "X-Phabricator-Csrf") {
				return h, true
			}
			return "", false
		}),
		gwrt.WithOutgoingHeaderMatcher(func(h string) (string, bool) {
			if stringz.PrefixFold(h, "access-control-allow-") ||
				strings.EqualFold(h, "Set-Cookie") ||
				strings.EqualFold(h, "vary") ||
				strings.EqualFold(h, "x-content-type-options") ||
				stringz.PrefixFold(h, "x-ratelimit-") {
				return h, true
			}
			return "", false
		}),
		gwrt.WithMetadata(func(c context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs("x-forwarded-method", req.Method)
		}),
		gwrt.WithProtoErrorHandler(gwrt.DefaultHTTPProtoErrorHandler),
	)
	var grpcDialOptions []grpc.DialOption
	if useTLS {
		grpcDialOptions = []grpc.DialOption{
			grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, s.APIDomain)),
		}
	} else {
		grpcDialOptions = []grpc.DialOption{grpc.WithInsecure()}
	}
	if s.EnableCORS {
		endpoints.ProxyServerCorsPattern.RegisterHandler(gwMux, s.CORSOriginHost, s.CORSOriginAllowSubdomain)
	}

	addr := l.Addr().String()
	addr = "127.0.0.1" + addr[strings.LastIndex(addr, ":"):]
	for _, endpoint := range s.ProxyEndpoints {
		log.Infoln("Registering endpoint:", funcName(endpoint.RegisterFunc))
		endpoints.RegisterProxy(endpoint.RegisterFunc, context.Background(), gwMux, addr, grpcDialOptions)
	}
	return gwMux
}

func (s *apiServer) ServeHTTPS() {
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
	gwMux := s.newGatewayMux(l, true)

	srv := &http.Server{
		Addr:         s.SecureAddr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
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

func Run(cfg *options.Config) {
	cfgBytes, _ := json.MarshalIndent(cfg, " ", " ")
	log.Infoln("Configuration:", string(cfgBytes))

	http.Handle("/metrics", promhttp.Handler())

	apisPublic := &apiServer{
		SecureAddr:               cfg.SecureAddr,
		PlaintextAddr:            cfg.PlaintextAddr,
		APIDomain:                cfg.APIDomain,
		CACertFile:               cfg.CACertFile,
		CertFile:                 cfg.CertFile,
		KeyFile:                  cfg.KeyFile,
		EnableJavaClient:         cfg.EnableJavaClient,
		EnableCORS:               cfg.EnableCORS,
		CORSOriginHost:           cfg.CORSOriginHost,
		CORSOriginAllowSubdomain: cfg.CORSOriginAllowSubdomain,
		GRPCEndpoints:            endpoints.GRPCServerEndpoints,
		ProxyEndpoints:           endpoints.ProxyServerEndpoints,
	}
	go apisPublic.ListenAndServe()

	// Run Monitoring Server with both /metric and /debug
	go func() {
		if cfg.MonitoringServerAddr != "" {
			if err := http.ListenAndServe(cfg.MonitoringServerAddr, nil); err != nil {
				log.Errorln("Failed to start monitoring server, cause", err)
			}
		}
	}()
}

func funcName(i interface{}) string {
	name := gort.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}
