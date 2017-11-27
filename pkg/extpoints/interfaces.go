package extpoints

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	rls "k8s.io/helm/pkg/proto/hapi/services"
)

type Connector interface {
	UID() string
	Connect(context.Context) (*grpc.ClientConn, rls.ReleaseServiceClient, error)
}
