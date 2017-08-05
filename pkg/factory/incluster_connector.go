package factory

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
	addr, err := c.getTillerAddr(client, c.namespace())
	if err != nil {
		return nil, err
	}
	conn, err := Connect(addr)
	if err != nil {
		return nil, err
	}
	return rls.NewReleaseServiceClient(conn), nil
}

func (c *InClusterConnector) getTillerAddr(client clientset.Interface, tillerNamespace string) (string, error) {
	svcList, err := client.CoreV1().Services(tillerNamespace).List(metav1.ListOptions{
		LabelSelector: tillerLabelSelector,
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", svcList.Items[0], defaultTillerPort), nil
}

func (c *InClusterConnector) namespace() string {
	if ns := os.Getenv("OPERATOR_NAMESPACE"); ns != "" {
		return ns
	}
	if data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns
		}
	}
	return apiv1.NamespaceDefault
}
