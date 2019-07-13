package connectors

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	kutil "kmodules.xyz/client-go"
	"kmodules.xyz/client-go/meta"
	"kubepack.dev/swift/pkg/extpoints"
)

type InClusterConnector struct {
	cfg       Config
	namespace string
}

var _ extpoints.Connector = &InClusterConnector{}

const (
	UIDInClusterConnector = "incluster"
)

func NewInClusterConnector(cfg Config) extpoints.Connector {
	return &InClusterConnector{
		cfg:       cfg,
		namespace: meta.Namespace(),
	}
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
	svc, err := c.findTillerService(client, c.namespace)
	if err == kutil.ErrNotFound {
		svc, err = c.findTillerService(client, core.NamespaceAll)
	}
	if err != nil {
		return "", errors.Wrap(err, "failed to detect tiller address")
	}

	if svc.Namespace == c.namespace {
		return fmt.Sprintf("%s:%d", svc.Name, defaultTillerPort), nil
	}
	return fmt.Sprintf("%s.%s.svc:%d", svc.Name, svc.Namespace, defaultTillerPort), nil
}

func (c *InClusterConnector) findTillerService(client clientset.Interface, namespace string) (*core.Service, error) {
	svcList, err := client.CoreV1().Services(namespace).List(metav1.ListOptions{
		LabelSelector: tillerLabelSelector,
	})
	if err != nil {
		return nil, err
	}
	if len(svcList.Items) == 0 {
		return nil, kutil.ErrNotFound
	}
	if len(svcList.Items) > 1 {
		ids := make([]string, len(svcList.Items))
		for i, svc := range svcList.Items {
			ids[i] = svc.Namespace + "/" + svc.Name
		}
		return nil, errors.Errorf("multiple tiller services found: %s", strings.Join(ids, ", "))
	}
	return &svcList.Items[0], nil
}
