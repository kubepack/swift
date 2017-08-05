package factory

import (
	"time"

	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	defaultTillerPort = 44134
)

var (
	tillerLabelSelector = labels.SelectorFromSet(labels.Set{
		"app":  "helm",
		"name": "tiller",
	}).String()
)

// connect returns a grpc connection to tiller or error. The grpc dial options
// are constructed here.
func connect(target string) (conn *grpc.ClientConn, err error) {
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

	if conn, err = grpc.Dial(target, opts...); err != nil {
		return nil, err
	}
	return conn, nil
}
