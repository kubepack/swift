package server

import (
	"time"

	"github.com/spf13/pflag"
)

type TillerOptions struct {
	Connector            string // incluster, direct, kubeconfig, appscode
	TillerEndpoint       string
	CACertFile           string
	ClientCertFile       string
	ClientPrivateKeyFile string
	InsecureSkipVerify   bool
	Timeout              time.Duration
	KubeContext          string
}

func NewTillerOptions() *TillerOptions {
	return &TillerOptions{
		Timeout: 5 * time.Minute,
	}
}

func (opt *TillerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&opt.Connector, "connector", opt.Connector, "Name of connector used to connect to Tiller server. Valid values are: incluster, direct, kubeconfig, appscode")
	fs.StringVar(&opt.TillerEndpoint, "tiller-endpoint", opt.TillerEndpoint, "Endpoint of Tiller server, eg, [scheme://]host:port")
	fs.StringVar(&opt.CACertFile, "tiller-ca-file", opt.CACertFile, "File containing CA certificate for Tiller server")
	fs.StringVar(&opt.ClientCertFile, "tiller-client-cert-file", opt.ClientCertFile, "File container client TLS certificate for Tiller server")
	fs.StringVar(&opt.ClientPrivateKeyFile, "tiller-client-private-key-file", opt.ClientPrivateKeyFile, "File containing client TLS private key for Tiller server")
	fs.BoolVar(&opt.InsecureSkipVerify, "tiller-insecure-skip-verify", opt.InsecureSkipVerify, "Skip certificate verification for Tiller server")
	fs.DurationVar(&opt.Timeout, "tiller-timeout", opt.Timeout, "Timeout used to connect to Tiller server")
	fs.StringVar(&opt.KubeContext, "kube-context", opt.KubeContext, "Kube context used by 'kubeconfig' connection")
}
