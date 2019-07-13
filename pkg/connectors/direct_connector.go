package connectors

import (
	"golang.org/x/net/context"
	"kubepack.dev/swift/pkg/extpoints"
)

type DirectConnector struct {
	cfg Config
}

var _ extpoints.Connector = &DirectConnector{}

const (
	UIDDirectConnector = "direct"
)

func NewDirectConnector(cfg Config) extpoints.Connector {
	return &DirectConnector{cfg: cfg}
}

func (c *DirectConnector) UID() string {
	return UIDDirectConnector
}

func (c *DirectConnector) Connect(ctx context.Context) (context.Context, error) {
	conn, err := Connect(c.cfg)
	if err != nil {
		return ctx, err
	}
	return WithConnection(ctx, conn), nil
}

func (c *DirectConnector) Close(ctx context.Context) error {
	return Connection(ctx).Close()
}
