package factory

import (
	"errors"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type TunnelConnector struct {
}

func (s *TunnelConnector) GetTillerAddr(client clientset.Interface, config *rest.Config) (string, error) {
	podList, err := client.CoreV1().Pods(apiv1.NamespaceAll).List(metav1.ListOptions{
		LabelSelector: tillerLabelSelector,
	})
	if err != nil {
		return "", err
	}
	if len(podList.Items) == 0 {
		return "", errors.New("no tiller pod(s) found")
	}

	tunnel := newTunnel(client.CoreV1().RESTClient(), config, podList.Items[0].Namespace, podList.Items[0].Name, defaultTillerPort)
	if err := tunnel.forwardPort(); err != nil {
		return "", err
	}
	return fmt.Sprintf("127.0.0.1:%d", tunnel.Local), nil
}
