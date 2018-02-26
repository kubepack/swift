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

	cfg Config
}

var _ extpoints.Connector = &KubeconfigConnector{}

const (
	UIDKubeconfigConnector = "kubeconfig"
)

func NewKubeconfigConnector(cfg Config) extpoints.Connector {
	return &KubeconfigConnector{cfg: cfg}
}

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

	cfgCopy := c.cfg
	cfgCopy.Endpoint = fmt.Sprintf("127.0.0.1:%d", tunnel.Local)
	conn, err := Connect(cfgCopy)
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
		CurrentContext:  c.cfg.KubeContext,
		ClusterDefaults: clientcmd.ClusterDefaults,
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, overrides).ClientConfig()
}
