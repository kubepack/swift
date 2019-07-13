package options

import (
	"flag"

	"github.com/spf13/pflag"
	"gomodules.xyz/grpc-go-addons/server"
)

type SecureServingOptions struct {
	SecureAddr    string
	PlaintextAddr string
	APIDomain     string
	CACertFile    string
	CertFile      string
	KeyFile       string
}

func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		SecureAddr:    ":8443",
		PlaintextAddr: ":8080",
	}
}

func (o *SecureServingOptions) AddGoFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.SecureAddr, "secure-addr", o.SecureAddr, "host:port used to serve secure apis")
	fs.StringVar(&o.PlaintextAddr, "plaintext-addr", o.PlaintextAddr, "host:port used to serve http json apis")

	fs.StringVar(&o.APIDomain, "api-domain", o.APIDomain, "Domain used for apiserver (prod: api.appscode.com")
	fs.StringVar(&o.CACertFile, "tls-ca-file", o.CACertFile, "File containing CA certificate")
	fs.StringVar(&o.CertFile, "tls-cert-file", o.CertFile, "File container server TLS certificate")
	fs.StringVar(&o.KeyFile, "tls-private-key-file", o.KeyFile, "File containing server TLS private key")
}

func (o *SecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	gfs := flag.NewFlagSet("grpc-serving", flag.ExitOnError)
	o.AddGoFlags(gfs)
	fs.AddGoFlagSet(gfs)
}

func (o *SecureServingOptions) ApplyTo(cfg *server.Config) error {
	cfg.SecureAddr = o.SecureAddr
	cfg.PlaintextAddr = o.PlaintextAddr
	cfg.APIDomain = o.APIDomain
	cfg.CACertFile = o.CACertFile
	cfg.CertFile = o.CertFile
	cfg.KeyFile = o.KeyFile

	return nil
}

func (o *SecureServingOptions) Validate() []error {
	return nil
}
