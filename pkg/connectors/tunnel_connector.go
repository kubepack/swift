package connectors

import (
	"errors"

	"github.com/appscode/kutil/tools/portforward"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type TunnelConnector struct {
}

func (s *TunnelConnector) GetTillerAddr(client clientset.Interface, config *rest.Config) (*portforward.Tunnel, error) {
	podList, err := client.CoreV1().Pods(core.NamespaceAll).List(metav1.ListOptions{
		LabelSelector: tillerLabelSelector,
	})
	if err != nil {
		return nil, err
	}
	if len(podList.Items) == 0 {
		return nil, errors.New("no tiller pod(s) found")
	}

	tunnel := portforward.NewTunnel(client.CoreV1().RESTClient(), config, podList.Items[0].Namespace, podList.Items[0].Name, defaultTillerPort)
	if err := tunnel.ForwardPort(); err != nil {
		return nil, err
	}
	return tunnel, nil
}
