package interceptors

import (
	"github.com/appscode/go/container/serializer"
	"github.com/appscode/go/log"
	"github.com/appscode/swift/pkg/server/endpoints"
	"golang.org/x/net/context"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewUnaryInterceptor(enableCORS bool, allowHost string, allowSubdomain bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if enableCORS {
			endpoints.SetCORSHeaders(ctx, allowHost, allowSubdomain)
		}
		endpoints.SetSecurityHeaders(ctx)

		defer func() {
			if r := recover(); r != nil {
				s := &spb.Status{
					Code:    int32(codes.Internal),
					Message: "Server crashed, :(",
				}
				err = status.FromProto(s).Err()
				return
			}
		}()

		var handlers = serializer.New()
		handlers.Add(&MonitorInterceptor{})
		// ref: https://github.com/mwitkow/go-grpc-middleware/blob/master/chain.go#L17
		buildChain := func(current grpc.UnaryServerInterceptor, next grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return current(currentCtx, currentReq, info, next)
			}
		}
		chain := handler
		for it := handlers.Iterator(); it.HasNext(); {
			n := it.Now()
			val, ok := n.(Interceptor)
			if !ok {
				log.Errorln("Bad Interceptor Registered")
				continue
			}
			chain = buildChain(val.Intercept, chain)
		}
		return chain(ctx, req)
	}
}
