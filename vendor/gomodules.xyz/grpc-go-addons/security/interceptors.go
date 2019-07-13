package security

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor returns a new unary server interceptor that add security header
//
// Invalid messages will be rejected with `Internal` before reaching any userspace handlers.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := setSecurityHeaders(ctx); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new unary server interceptor that add security header
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := setSecurityHeaders(stream.Context())
		if err != nil {
			return err
		}
		return handler(srv, stream)
	}
}

func setSecurityHeaders(ctx context.Context) error {
	headers := map[string]string{
		"x-content-type-options": "nosniff", // http://stackoverflow.com/a/3146618/244009
	}
	return grpc.SetHeader(ctx, metadata.New(headers))
}
