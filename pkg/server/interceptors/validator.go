package interceptors

import (
	"github.com/appscode/api"
	"github.com/appscode/api/dtypes"
	"github.com/appscode/errors"
	"github.com/appscode/go/arrays"
	"github.com/appscode/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ValidationInterceptor struct {
	Order        int      `json:"order,omitempty"`
	NoValidation []string `json:"noValidation,omitempty"`
}

func (v *ValidationInterceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var causes string = "request contains invalid data"
	if found, _ := arrays.Contains(v.NoValidation, info.FullMethod); !found {
		log.V(5).Infoln("validating request for", info.FullMethod)
		requestInterface, ok := req.(api.Request)
		invalid := true
		if ok {
			validator, err := requestInterface.IsValid()
			if err == nil {
				invalid = !validator.Valid()
				log.Debugln("validator logs", validator.Errors())
				if c := validator.Errors(); len(c) > 0 {
					causes = c[0].Field() + ": " + c[0].Description()
				}
			}
		}

		if invalid {
			return nil, dtypes.InvalidArgument(errors.New(causes).Err())
		}
	}
	return handler(ctx, req)
}

func (v *ValidationInterceptor) Weight() int {
	return v.Order
}
