package endpoints

import (
	"reflect"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	Register interface{}
	Server   interface{}
}

type GRPCRegistry []*grpcHandler

func (r *GRPCRegistry) Register(fn, server interface{}) {
	if *r == nil {
		*r = make([]*grpcHandler, 0)
	}
	*r = append(*r, &grpcHandler{
		Register: fn,
		Server:   server,
	})
}

func (r GRPCRegistry) ApplyTo(srv *grpc.Server) error {
	for _, ep := range r {
		glog.Infof("Registering grpc server: %s", reflect.TypeOf(ep.Server))

		fn := reflect.ValueOf(ep.Register)
		params := []reflect.Value{
			reflect.ValueOf(srv),
			reflect.ValueOf(ep.Server),
		}
		fn.Call(params)
	}
	return nil
}
