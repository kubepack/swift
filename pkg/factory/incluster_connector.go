package factory

import (
	"errors"
	"fmt"
	"strings"

	"github.com/appscode/swift/pkg/extpoints"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

func (c *InClusterConnector) Connect(ctx context.Context) (*grpc.ClientConn, rls.ReleaseServiceClient, error) {
	config, err := restclient.InClusterConfig()
	if err != nil {
		return nil, nil, err
	}
	client, err := clientset.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	addr, err := c.getTillerAddr(client)
	if err != nil {
		return nil, nil, err
	}
	conn, err := Connect(addr)
	if err != nil {
		return nil, nil, err
	}
	return conn, rls.NewReleaseServiceClient(conn), nil
}

func (c *InClusterConnector) getTillerAddr(client clientset.Interface) (string, error) {
	svcList, err := client.CoreV1().Services(apiv1.NamespaceAll).List(metav1.ListOptions{
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
