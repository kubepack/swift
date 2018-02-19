package connectors

import (
	"github.com/appscode/swift/pkg/extpoints"
	"golang.org/x/net/context"
)

type DirectConnector struct {
	TillerEndpoint       string
	TillerCACertFile     string
	TillerClientCertFile string
	TillerClientKeyFile  string
	InsecureSkipVerify   bool
}

var _ extpoints.Connector = &DirectConnector{}

const (
	UIDDirectConnector = "direct"
)

func (c *DirectConnector) UID() string {
	return UIDDirectConnector
}

func (c *DirectConnector) Connect(ctx context.Context) (context.Context, error) {
	conn, err := Connect(c.TillerEndpoint, c.TillerCACertFile, c.TillerClientCertFile, c.TillerClientKeyFile, c.InsecureSkipVerify)
	if err != nil {
		return ctx, err
	}
	return WithConnection(ctx, conn), nil
}

func (c *DirectConnector) Close(ctx context.Context) error {
	return Connection(ctx).Close()
}
