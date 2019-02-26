package connectors

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	grpc_glog "github.com/grpc-ecosystem/go-grpc-middleware/logging/glog"
	ctx_glog "github.com/grpc-ecosystem/go-grpc-middleware/tags/glog"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	defaultTillerPort = 44134
	// maxReceiveMsgSize uses 20MB as the default message size limit.
	// the gRPC's default size is 4MB.
	// Since Tiller has been change the messages' size to 20MB, so we should make this value to 20MB.
	maxReceiveMsgSize = 1024 * 1024 * 20
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
	optsGLog := []grpc_glog.Option{
		grpc_glog.WithDecider(func(methodFullName string, err error) bool {
			return cfg.LogRPC
		}),
	}
	glogEntry := ctx_glog.NewEntry(ctx_glog.Logger)
	opts := []grpc.DialOption{
		grpc.WithBlock(), // required for timeout
		grpc.WithUnaryInterceptor(grpc_glog.UnaryClientInterceptor(glogEntry, optsGLog...)),
		grpc.WithStreamInterceptor(grpc_glog.StreamClientInterceptor(glogEntry, optsGLog...)),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxReceiveMsgSize)),
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
