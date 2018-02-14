package factory

import (
	"fmt"
	"time"

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
func Connect(target string, tillerCACertFile string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithBlock(), // required for timeout
	}
	if tillerCACertFile == "" {
		opts = append(opts, grpc.WithInsecure())
	} else {
		cred, err := credentials.NewClientTLSFromFile(tillerCACertFile, "")
		if err != nil {
			return nil, fmt.Errorf("failed load tiller ca cert file. Reson: %s", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return grpc.DialContext(ctx, target, opts...)
}
