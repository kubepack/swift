package factory

import (
	"github.com/appscode/wheel/pkg/extpoints"
	"golang.org/x/net/context"
	rls "k8s.io/helm/pkg/proto/hapi/services"
)

type DirectConnector struct {
	TillerEndpoint string
}

var _ extpoints.Connector = &DirectConnector{}

const (
	UIDDirectConnector = "direct"
)

func (c *DirectConnector) UID() string {
	return UIDDirectConnector
}

func (c *DirectConnector) Connect(ctx context.Context) (rls.ReleaseServiceClient, error) {
	conn, err := Connect(c.TillerEndpoint)
	if err != nil {
		return nil, err
	}
	return rls.NewReleaseServiceClient(conn), nil
}
