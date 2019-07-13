package cors

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
// https://fetch.spec.whatwg.org/#cors-protocol-and-credentials
// For requests without credentials, the server may specify "*" as a wildcard, thereby allowing any origin to access the resource.
type PatternRegistry []runtime.Pattern

func (r *PatternRegistry) Register(f []runtime.Pattern) {
	if *r == nil {
		*r = make([]runtime.Pattern, 0)
	}
	*r = append(*r, f...)
}
