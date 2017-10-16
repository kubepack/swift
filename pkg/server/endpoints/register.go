package endpoints

import (
	"reflect"
	"sync"

	"github.com/appscode/go/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type endPoint struct {
	RegisterFunc interface{}
	Server       interface{}
}

func RegisterGRPC(function interface{}, s *grpc.Server, val interface{}) {
	f := reflect.ValueOf(function)
	var mu sync.Mutex

	mu.Lock()
	params := []reflect.Value{
		reflect.ValueOf(s),
		reflect.ValueOf(val),
	}
	f.Call(params)
	mu.Unlock()
}

func RegisterProxy(function interface{}, ctx context.Context, mux *runtime.ServeMux, url string, opts []grpc.DialOption) {
	f := reflect.ValueOf(function)
	var mu sync.Mutex

	mu.Lock()
	params := []reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(mux),
		reflect.ValueOf(url),
		reflect.ValueOf(opts),
	}
	vals := f.Call(params)
	mu.Unlock()
	if !vals[0].IsNil() {
		// checking if any error exists
		log.Fatalln("proxy can not connect to server")
	}
}
