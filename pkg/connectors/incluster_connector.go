package connectors

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/appscode/swift/pkg/extpoints"
	"golang.org/x/net/context"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

type InClusterConnector struct {
	TillerCACertFile     string
	TillerClientCertFile string
	TillerClientKeyFile  string
	InsecureSkipVerify   bool
	Timeout              time.Duration
}

var _ extpoints.Connector = &InClusterConnector{}

const (
	UIDInClusterConnector = "incluster"
)

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
	addr, err := c.getTillerAddr(client)
	if err != nil {
		return ctx, err
	}
	conn, err := Connect(addr, c.TillerCACertFile, c.TillerClientCertFile, c.TillerClientKeyFile, c.InsecureSkipVerify, c.Timeout)
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
		return "", fmt.Errorf("multiple tiller services found: %s", strings.Join(ids, ", "))
	}
	return fmt.Sprintf("%s.%s:%d", svcList.Items[0].Name, svcList.Items[0].Namespace, defaultTillerPort), nil
}
