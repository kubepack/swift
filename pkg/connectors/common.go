package connectors

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
func Connect(cfg Config) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithBlock(), // required for timeout
	}
	if cfg.InsecureSkipVerify {
		opts = append(opts, grpc.WithInsecure())
	} else {
		tlsConfig := &tls.Config{}

		// load cacert
		if cfg.CACertFile != "" {
			caCert, err := ioutil.ReadFile(cfg.CACertFile)
			if err != nil {
				return nil, errors.Wrap(err, "failed to load ca cert")
				return nil, err
			}
			pool := x509.NewCertPool()
			pool.AppendCertsFromPEM(caCert)
			tlsConfig.RootCAs = pool
		}

		// load client cert/key
		if cfg.ClientCertFile != "" && cfg.ClientPrivateKeyFile != "" {
			pair, err := tls.LoadX509KeyPair(cfg.ClientCertFile, cfg.ClientPrivateKeyFile)
			if err != nil {
				return nil, errors.Wrap(err, "load client cert/key.")
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
		}

		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	}

	ctx, _ := context.WithTimeout(context.Background(), cfg.Timeout)
	return grpc.DialContext(ctx, cfg.Endpoint, opts...)
}
