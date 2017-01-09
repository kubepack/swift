package app

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	gort "runtime"
	"strings"
	"time"

	"github.com/appscode/go/runtime"
	"github.com/appscode/grpc-seed/cmd/seed-apis/app/options"
	"github.com/appscode/grpc-seed/pkg/apiserver/endpoints"
	"github.com/appscode/log"
	gwrt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type apiServer struct {
	Port       int
	CACertFile string
	CertFile   string
	KeyFile    string
}

func (s *apiServer) UseTLS() bool {
	return !(s.CACertFile == "" && s.CertFile == "" && s.KeyFile == "")
}

func (s *apiServer) ListenAndServe() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)
	if !s.UseTLS() {
		grpcl := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
		httpl := m.Match(cmux.Any())

		go s.ServeGRPC(grpcl)
		go s.ServeHTTP(httpl)
	} else {
		grpcl := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
		tlsl := m.Match(cmux.Any())

		go s.ServeGRPC(grpcl)
		go s.RedirectToHTTPS()
		go s.ServeHTTPS(tlsl)
	}

	log.Fatalln(m.Serve())
}

func (s *apiServer) ServeGRPC(l net.Listener) {
	defer runtime.HandleCrash()

	var gRPCServer *grpc.Server
	if s.UseTLS() {
		creds, err := credentials.NewServerTLSFromFile(s.CertFile, s.KeyFile)
		if err != nil {
			log.Fatalln(err)
		}
		gRPCServer = grpc.NewServer(grpc.Creds(creds))
	} else {
		gRPCServer = grpc.NewServer()
	}

	for _, endpoint := range endpoints.GRPCServerEndpoints {
		log.Infoln("Registering server:", reflect.TypeOf(endpoint.Server))
		endpoints.RegisterGRPC(endpoint.RegisterFunc, gRPCServer, endpoint.Server)
	}

	log.Infoln("[GRPCSERVER] Starting gRPC Server at port", s.Port)
	log.Fatalln("[GRPCSERVER] gRPC Server failed:", gRPCServer.Serve(l))
}

func (s *apiServer) ServeHTTP(l net.Listener) {
	defer runtime.HandleCrash()

	gwMux := gwrt.NewServeMux()
	var grpcDialOptions []grpc.DialOption
	if s.UseTLS() {
		grpcDialOptions = []grpc.DialOption{
			grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "appscode.stream")),
		}
	} else {
		grpcDialOptions = []grpc.DialOption{grpc.WithInsecure()}
	}
	for _, endpoint := range endpoints.ProxyServerEndpoints {
		log.Infoln("Registering endpoint:", funcName(endpoint.RegisterFunc))
		endpoints.RegisterProxy(endpoint.RegisterFunc, context.Background(), gwMux, fmt.Sprintf("127.0.0.1:%d", s.Port), grpcDialOptions)
	}

	log.Infoln("[PROXYSERVER] Sarting Proxy Server at port", s.Port)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.Port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      gwMux,
	}
	log.Fatalln("[PROXYSERVER] Proxy Server failed:", srv.Serve(l))
}

func (s *apiServer) RedirectToHTTPS() {
	defer runtime.HandleCrash()
	log.Infoln("[REDIRECTOR] Sarting Redirector Server")
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Connection", "close")
			url := "https://" + req.Host + req.URL.String()
			http.Redirect(w, req, url, http.StatusMovedPermanently)
		}),
	}
	log.Fatalln("[REDIRECTOR] Redirector Server failed:", srv.ListenAndServe())
}

func (s *apiServer) ServeHTTPS(l net.Listener) {
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
			// tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			// tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		ClientAuth: tls.VerifyClientCertIfGiven,
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

	// Create TLS listener.
	tlsl := tls.NewListener(l, tlsConfig)

	// Serve HTTP over TLS.
	s.ServeHTTP(tlsl)
}

func Run(cfg *options.Config) {
	cfgBytes, _ := json.Marshal(cfg)
	log.Infoln("Configuration:", string(cfgBytes))

	server := &apiServer{
		Port:       cfg.APIPort,
		CACertFile: cfg.CACertFile,
		CertFile:   cfg.CertFile,
		KeyFile:    cfg.KeyFile,
	}
	go server.ListenAndServe()

	go func() {
		log.Infoln(http.ListenAndServe(fmt.Sprintf(":%d", cfg.PprofPort), nil))
	}()
}

func funcName(i interface{}) string {
	name := gort.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}
