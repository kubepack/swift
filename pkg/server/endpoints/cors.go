package endpoints

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	_env "github.com/appscode/go/env"
	"github.com/appscode/go/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
// https://fetch.spec.whatwg.org/#cors-protocol-and-credentials
// For requests without credentials, the server may specify "*" as a wildcard, thereby allowing any origin to access the resource.
type corsPattern struct {
	patterns []runtime.Pattern
	origin   string
}

func (p *corsPattern) Register(f []runtime.Pattern) {
	if p.patterns == nil {
		p.patterns = make([]runtime.Pattern, 0)
	}
	p.patterns = append(p.patterns, f...)
}

func (p *corsPattern) RegisterHandler(mux *runtime.ServeMux, allowHost string, allowSubdomain bool) {
	for _, p := range p.patterns {
		mux.Handle("OPTIONS", p, newCORSHandler(allowHost, allowSubdomain))
	}
}

var ProxyServerCorsPattern = &corsPattern{}

func newCORSHandler(allowHost string, allowSubdomain bool) runtime.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
		headers := map[string]string{
			"access-control-allow-methods": "POST,GET,OPTIONS,PUT,DELETE",
			"access-control-allow-headers": req.Header.Get("access-control-request-headers"),
		}
		if allowHost == "*" {
			headers["access-control-allow-origin"] = "*"
		} else if allowHost != "" {
			origin := req.Header.Get("Origin")

			u, err := url.Parse(origin)
			if err != nil {
				log.Errorln("Failed to parse CORS origin header", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ok := u.Host == allowHost ||
				(allowSubdomain && strings.HasSuffix(u.Host, "."+allowHost))
			if !ok {

				log.Errorln("CORS request from prohibited domain %v", origin)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !_env.FromHost().DevMode() {
				u.Scheme = "https"
			}
			headers["access-control-allow-origin"] = u.String()
			headers["access-control-allow-credentials"] = "true"
			headers["vary"] = "Origin"
		}
		for k, v := range headers {
			w.Header().Set(k, v)
		}
	}
}

func SetCORSHeaders(ctx context.Context, allowHost string, allowSubdomain bool) error {
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
	if allowHost == "*" {
		headers["access-control-allow-origin"] = "*"
	} else if allowHost != "" {
		var origin string
		if origins, ok := md["origin"]; ok {
			origin = origins[0]
		}

		u, err := url.Parse(origin)
		if err != nil {
			return errors.New("Failed to parse CORS origin header")
		}
		ok := u.Host == allowHost ||
			(allowSubdomain && strings.HasSuffix(u.Host, "."+allowHost))
		if !ok {
			return fmt.Errorf("CORS request from prohibited domain %v", origin)
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
