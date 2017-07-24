package app

import (
	"fmt"

	app "github.com/appscode/grpc-seed/pkg/apis/app/v1beta1"
	"github.com/appscode/grpc-seed/pkg/apiserver/endpoints"
	"golang.org/x/net/context"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(app.RegisterAppsServer, &AppsServer{})
	endpoints.ProxyServerEndpoints.Register(app.RegisterAppsHandlerFromEndpoint)
}

type AppsServer struct{}

var _ app.AppsServer = &AppsServer{}

func (*AppsServer) Hello(ctx context.Context, req *app.HelloRequest) (*app.HelloResponse, error) {
	return &app.HelloResponse{
		Greetings: fmt.Sprintf("Hello, %s.", req.Name),
	}, nil
}
