package app

import (
	"fmt"
	"time"

	"github.com/appscode/lever/pkg/kube"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	rls "k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/version"
)

const (
	DEFAULT_OPERATOR_NS   = "kube-system"
	OPERATOR_NS_HEADER    = "k8s-operator-ns"
	DEFAULT_OPERATOR_PORT = 44134
	LOCALHOST             = "127.0.0.1"
	HELM_APP_LABEL        = "helm"
)

// connect returns a grpc connection to tiller or error. The grpc dial options
// are constructed here.
func connect(host string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTimeout(5 * time.Second),
		grpc.WithBlock(),
	}
	opts = append(opts, grpc.WithInsecure())

	/*switch {
	case h.opts.useTLS:
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(h.opts.tlsConfig)))
	default:
		opts = append(opts, grpc.WithInsecure())
	}*/

	if conn, err = grpc.Dial(host, opts...); err != nil {
		return nil, err
	}
	return conn, nil
}

// NewContext creates a versioned context.
func newContext() context.Context {
	md := metadata.Pairs("x-helm-api-client", version.GetVersion())
	return metadata.NewContext(context.TODO(), md)
}

func getHost(operatorNamespace string) (string, error) {
	if operatorNamespace == "" {
		operatorNamespace = DEFAULT_OPERATOR_NS
	}
	operatorPortNumber := DEFAULT_OPERATOR_PORT

	f := kube.NewKubeFactory()
	restClient, err := f.RESTClient()
	if err != nil {
		return "", err
	}

	config, err := f.ClientConfig()
	if err != nil {
		return "", err
	}

	clientSet, err := f.ClientSet()
	if err != nil {
		return "", err
	}

	operatorPodList, err := clientSet.Core().Pods(operatorNamespace).List(metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(labels.Set{
			"app": HELM_APP_LABEL,
		}).String(),
	})
	if err != nil {
		return "", err
	}

	tunnel := newTunnel(restClient, config, operatorNamespace, operatorPodList.Items[0].Name, operatorPortNumber)
	if err := tunnel.forwardPort(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%d", LOCALHOST, tunnel.Local), nil
}

func getReleaseServiceClient(ctx context.Context) (rls.ReleaseServiceClient, error) {
	var operatorNs string

	if headers, ok := metadata.FromContext(ctx); ok {
		token := headers[OPERATOR_NS_HEADER]
		if len(token) > 0 {
			operatorNs = token[0]
		}
	}

	host, err := getHost(operatorNs)
	if err != nil {
		return nil, err
	}

	conn, err := connect(host)
	if err != nil {
		return nil, err
	}

	return rls.NewReleaseServiceClient(conn), nil
}
