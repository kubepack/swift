package factory

import (
	"fmt"

	"github.com/appscode/wheel/pkg/extpoints"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	restclient "k8s.io/client-go/rest"
	rls "k8s.io/helm/pkg/proto/hapi/services"
)

type InClusterConnector struct {
}

var _ extpoints.Connector = &InClusterConnector{}

const (
	UIDInClusterConnector = "incluster"
)

func (c *InClusterConnector) UID() string {
	return UIDInClusterConnector
}

func (c *InClusterConnector) Connect(ctx context.Context) (rls.ReleaseServiceClient, error) {
	config, err := restclient.InClusterConfig()
	if err != nil {
		return nil, err
	}
	client, err := clientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	addr, err := c.getTillerAddr(client)
	if err != nil {
		return nil, err
	}
	conn, err := Connect(addr)
	if err != nil {
		return nil, err
	}
	return rls.NewReleaseServiceClient(conn), nil
}

func (c *InClusterConnector) getTillerAddr(client clientset.Interface) (string, error) {
	svcList, err := client.CoreV1().Services(apiv1.NamespaceAll).List(metav1.ListOptions{
		LabelSelector: tillerLabelSelector,
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s:%d", svcList.Items[0].Name, svcList.Items[0].Namespace, defaultTillerPort), nil
}
