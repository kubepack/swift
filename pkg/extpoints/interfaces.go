package extpoints

import (
	"golang.org/x/net/context"
	rls "k8s.io/helm/pkg/proto/hapi/services"
)

type Connector interface {
	UID() string
	Connect(context.Context) (rls.ReleaseServiceClient, error)
}
