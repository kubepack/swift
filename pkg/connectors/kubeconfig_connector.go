package connectors

import (
	"fmt"

	"github.com/appscode/swift/pkg/extpoints"
	"golang.org/x/net/context"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeconfigConnector struct {
	*TunnelConnector

	Context            string
	InsecureSkipVerify bool
}

var _ extpoints.Connector = &KubeconfigConnector{}

const (
	UIDKubeconfigConnector = "kubeconfig"
)

func (c *KubeconfigConnector) UID() string {
	return UIDKubeconfigConnector
}

func (c *KubeconfigConnector) Connect(ctx context.Context) (context.Context, error) {
	config, err := c.getConfig()
	if err != nil {
		return ctx, err
	}
	client, err := clientset.NewForConfig(config)
	if err != nil {
		return ctx, err
	}
	tunnel, err := c.GetTillerAddr(client, config)
	if err != nil {
		return ctx, err
	}
	ctx = WithTunnel(ctx, tunnel)

	addr := fmt.Sprintf("127.0.0.1:%d", tunnel.Local)
	conn, err := Connect(addr, "", "", "", c.InsecureSkipVerify)
	if err != nil {
		return ctx, err
	}
	return WithConnection(ctx, conn), nil
}

func (c *KubeconfigConnector) Close(ctx context.Context) error {
	defer Tunnel(ctx).Close()
	return Connection(ctx).Close()
}

func (c *KubeconfigConnector) getConfig() (*rest.Config, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.DefaultClientConfig = &clientcmd.DefaultClientConfig

	overrides := &clientcmd.ConfigOverrides{
		CurrentContext:  c.Context,
		ClusterDefaults: clientcmd.ClusterDefaults,
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, overrides).ClientConfig()
}
