package app

import (
	"encoding/json"
	"fmt"
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
)

type apiServer struct {
	Port int

	GRPCServer *grpc.Server
	ProxyMux   *gwrt.ServeMux
}

func (s *apiServer) ListenAndServe() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)

	// We first match the connection against HTTP2 fields. If matched, the
	// connection will be sent through the "grpcl" listener.
	grpcl := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	// Otherwise, we match it againts HTTP1 methods. If matched,
	// it is sent through the "httpl" listener.
	httpl := m.Match(cmux.HTTP1Fast())

	// Then we used the muxed listeners.
	go s.ServeGRPC(grpcl)
	go s.ServeHTTP(httpl)

	log.Fatalln(m.Serve())
}

func (s *apiServer) ServeGRPC(l net.Listener) {
	defer runtime.HandleCrash()
	for _, endpoint := range endpoints.GRPCServerEndpoints {
		log.Infoln("Registering server:", reflect.TypeOf(endpoint.Server))
		endpoints.RegisterGRPC(endpoint.RegisterFunc, s.GRPCServer, endpoint.Server)
	}

	log.Infoln("[GRPCSERVER] Starting gRPC Server at port", s.Port)
	log.Fatalln("[GRPCSERVER] gRPC Server failed:", s.GRPCServer.Serve(l))
}

var grpcDialOptions = []grpc.DialOption{grpc.WithInsecure()}

func (s *apiServer) ServeHTTP(l net.Listener) {
	defer runtime.HandleCrash()
	for _, endpoint := range endpoints.ProxyServerEndpoints {
		log.Infoln("Registering endpoint:", funcName(endpoint.RegisterFunc))
		endpoints.RegisterProxy(endpoint.RegisterFunc, context.Background(), s.ProxyMux, fmt.Sprintf("127.0.0.1:%d", s.Port), grpcDialOptions)
	}

	log.Infoln("[PROXYSERVER] Sarting Proxy Server at port", s.Port)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      s.ProxyMux,
	}
	log.Fatalln("[PROXYSERVER] Proxy Server failed:", srv.Serve(l))
}

func Run(cfg *options.Config) {
	cfgBytes, _ := json.Marshal(cfg)
	log.Infoln("Configuration:", string(cfgBytes))

	server := &apiServer{
		Port:       cfg.APIPort,
		GRPCServer: grpc.NewServer(),
		ProxyMux:   gwrt.NewServeMux(),
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
