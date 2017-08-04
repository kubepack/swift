package interceptors

import (
	"fmt"
	"os"

	"github.com/appscode/api/meta"
	"github.com/appscode/go/encoding/yaml"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Interceptor interface {
	Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
	Weight() int
}

type InterceptorConfig struct {
	Type                      string                     `json:"type,omitempty"`
	ValidationInterceptor     *ValidationInterceptor     `json:"validationInterceptor,omitempty"`
}

// InterceptorConfigs needs to run before any other init runs in
// package.
var InterceptorConfigs = Load()

func Load() *InterceptorConfig {
	b, err := meta.Asset("meta/config.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	i := &InterceptorConfig{}
	err = yaml.Unmarshal(b, i)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return i
}
