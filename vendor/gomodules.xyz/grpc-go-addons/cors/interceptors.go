package cors

import (
	"net/url"
	"strings"

	_env "gomodules.xyz/x/env"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor returns a new unary server interceptor for OpenTracing.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := setCORSHeaders(ctx, o); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for OpenTracing.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := setCORSHeaders(stream.Context(), o)
		if err != nil {
			return err
		}
		return handler(srv, stream)
	}
}

func setCORSHeaders(ctx context.Context, opts *options) error {
	headers := map[string]string{
		"access-control-allow-methods": "POST,GET,OPTIONS,PUT,DELETE",
	}
	var md metadata.MD
	if m, ok := metadata.FromIncomingContext(ctx); ok {
		md = m
	}
	if rh, ok := md["access-control-request-headers"]; ok {
		headers["access-control-allow-headers"] = rh[0]
	}
	if opts.allowHost == "*" {
		headers["access-control-allow-origin"] = "*"
	} else if opts.allowHost != "" {
		var origin string
		if origins, ok := md["origin"]; ok {
			origin = origins[0]
		}

		u, err := url.Parse(origin)
		if err != nil {
			return errors.New("Failed to parse CORS origin header")
		}
		ok := u.Host == opts.allowHost ||
			(opts.allowSubdomain && strings.HasSuffix(u.Host, "."+opts.allowHost))
		if !ok {
			return errors.Errorf("CORS request from prohibited domain %v", origin)
		}
		if !_env.FromHost().DevMode() {
			u.Scheme = "https"
		}
		headers["access-control-allow-origin"] = u.String()
		headers["access-control-allow-credentials"] = "true"
		headers["vary"] = "Origin"
	}
	return grpc.SetHeader(ctx, metadata.New(headers))
}
