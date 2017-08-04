package endpoints

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GRPCEndpoints []*endPoint

func (s *GRPCEndpoints) Register(fun, server interface{}) {
	if *s == nil {
		*s = make([]*endPoint, 0)
	}
	*s = append(*s, &endPoint{
		RegisterFunc: fun,
		Server:       server,
	})
}

// all the public endpoints that will be exposed are listed
var GRPCServerEndpoints = GRPCEndpoints{}

func SetSecurityHeaders(ctx context.Context) error {
	headers := map[string]string{
		"x-content-type-options": "nosniff", // http://stackoverflow.com/a/3146618/244009
	}
	return grpc.SetHeader(ctx, metadata.New(headers))
}
