package interceptors

import (
	goprom "github.com/grpc-ecosystem/go-grpc-prometheus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type MonitorInterceptor struct{}

func (m *MonitorInterceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return goprom.UnaryServerInterceptor(ctx, req, info, handler)
}

func (m *MonitorInterceptor) Weight() int {
	return 0
}
