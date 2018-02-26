package connectors

import (
	"fmt"
	"strings"

	"github.com/appscode/swift/pkg/extpoints"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

type InClusterConnector struct {
	cfg Config
}

var _ extpoints.Connector = &InClusterConnector{}

const (
	UIDInClusterConnector = "incluster"
)

func NewInClusterConnector(cfg Config) extpoints.Connector {
	return &InClusterConnector{cfg: cfg}
}

func (c *InClusterConnector) UID() string {
	return UIDInClusterConnector
}

func (c *InClusterConnector) Connect(ctx context.Context) (context.Context, error) {
	config, err := restclient.InClusterConfig()
	if err != nil {
		return ctx, err
	}
	client, err := clientset.NewForConfig(config)
	if err != nil {
		return ctx, err
	}
	cfgCopy := c.cfg
	cfgCopy.Endpoint, err = c.getTillerAddr(client)
	if err != nil {
		return ctx, err
	}
	conn, err := Connect(cfgCopy)
	if err != nil {
		return ctx, err
	}
	return WithConnection(ctx, conn), nil
}

func (c *InClusterConnector) Close(ctx context.Context) error {
	return Connection(ctx).Close()
}

func (c *InClusterConnector) getTillerAddr(client clientset.Interface) (string, error) {
	svcList, err := client.CoreV1().Services(core.NamespaceAll).List(metav1.ListOptions{
		LabelSelector: tillerLabelSelector,
	})
	if err != nil {
		return "", err
	}
	if len(svcList.Items) == 0 {
		return "", errors.New("no tiller service found")
	}
	if len(svcList.Items) > 1 {
		ids := make([]string, len(svcList.Items))
		for i, svc := range svcList.Items {
			ids[i] = svc.Namespace + "/" + svc.Name
		}
		return "", errors.Errorf("multiple tiller services found: %s", strings.Join(ids, ", "))
	}
	return fmt.Sprintf("%s.%s:%d", svcList.Items[0].Name, svcList.Items[0].Namespace, defaultTillerPort), nil
}
