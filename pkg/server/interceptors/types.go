package interceptors

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Interceptor interface {
	Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
	Weight() int
}
