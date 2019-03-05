package connectors

import (
	"context"

	"google.golang.org/grpc"
	"kmodules.xyz/client-go/tools/portforward"
)

type paramConn struct{}
type paramTunnel struct{}

func Connection(ctx context.Context) *grpc.ClientConn {
	return ctx.Value(paramConn{}).(*grpc.ClientConn)
}

func WithConnection(parent context.Context, v *grpc.ClientConn) context.Context {
	return context.WithValue(parent, paramConn{}, v)
}

func Tunnel(ctx context.Context) *portforward.Tunnel {
	return ctx.Value(paramTunnel{}).(*portforward.Tunnel)
}

func WithTunnel(parent context.Context, v *portforward.Tunnel) context.Context {
	return context.WithValue(parent, paramTunnel{}, v)
}
