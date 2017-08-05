package factory

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	headerTillerNamespace = "k8s-tiller-namespace"
)

type TunnelConnector struct {
}

func (s *TunnelConnector) getTillerAddr(client clientset.Interface, config *rest.Config, tillerNamespace string) (string, error) {
	podList, err := client.CoreV1().Pods(tillerNamespace).List(metav1.ListOptions{
		LabelSelector: tillerLabelSelector,
	})
	if err != nil {
		return "", err
	}
	tunnel := newTunnel(client.CoreV1().RESTClient(), config, tillerNamespace, podList.Items[0].Name, defaultTillerPort)
	if err := tunnel.forwardPort(); err != nil {
		return "", err
	}
	return fmt.Sprintf("127.0.0.1:%d", tunnel.Local), nil
}

func (s *TunnelConnector) namespace(ctx context.Context) string {
	if headers, ok := metadata.FromContext(ctx); ok {
		namespaces := headers[headerTillerNamespace]
		if len(namespaces) > 0 {
			return namespaces[0]
		}
	}
	return "kube-system"
}
