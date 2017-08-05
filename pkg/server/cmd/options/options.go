package options

import (
	"github.com/spf13/pflag"
)

type Options struct {
	SecureAddr               string
	PlaintextAddr            string
	EnableJavaClient         bool
	APIDomain                string
	CACertFile               string
	CertFile                 string
	KeyFile                  string
	EnableCORS               bool
	CORSOriginHost           string
	CORSOriginAllowSubdomain bool
	WebAddr                  string
	EnableAnalytics          bool

	Connector      string // incluster, direct, kubeconfig, appscode
	TillerEndpoint string
	KubeContext    string
}

func New() *Options {
	return &Options{
		SecureAddr:      ":50055",
		PlaintextAddr:   ":9855",
		WebAddr:         ":56790",
		EnableAnalytics: true,
	}
}

func (opt *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&opt.SecureAddr, "secure-addr", opt.SecureAddr, "host:port used to server secure apis")
	fs.StringVar(&opt.PlaintextAddr, "plaintext-addr", opt.PlaintextAddr, "host:port used to server plaintext apis")

	fs.BoolVar(&opt.EnableJavaClient, "enable-java-client", opt.EnableJavaClient, "Set true to send SETTINGS frame from the server. Default set to false")
	fs.StringVar(&opt.APIDomain, "api-domain", opt.APIDomain, "Domain used to server wheel api")
	fs.StringVar(&opt.CACertFile, "caCertFile", opt.CACertFile, "File containing CA certificate")
	fs.StringVar(&opt.CertFile, "certFile", opt.CertFile, "File container server TLS certificate")
	fs.StringVar(&opt.KeyFile, "keyFile", opt.KeyFile, "File containing server TLS private key")

	fs.BoolVar(&opt.EnableCORS, "enable-cors", opt.EnableCORS, "Enable CORS support")
	fs.StringVar(&opt.CORSOriginHost, "cors-origin-host", opt.CORSOriginHost, `Allowed CORS origin host e.g, domain[:port]`)
	fs.BoolVar(&opt.CORSOriginAllowSubdomain, "cors-origin-allow-subdomain", opt.CORSOriginAllowSubdomain, "Allow CORS request from subdomains of origin")

	fs.StringVar(&opt.WebAddr, "web-addr", opt.WebAddr, "Address to listen on for web interface and telemetry.")
	fs.BoolVar(&opt.EnableAnalytics, "analytics", opt.EnableAnalytics, "Send analytical events to Google Analytics")

	fs.StringVar(&opt.Connector, "connector", opt.Connector, "Name of connector used to connect to Tiller server. Valid values are: incluster, direct, kubeconfig, appscode")
	fs.StringVar(&opt.TillerEndpoint, "tiller-endpoint", opt.TillerEndpoint, "Endpoint of Tiller server, eg, [scheme://]host:port")
	fs.StringVar(&opt.KubeContext, "kube-context", opt.KubeContext, "Kube context used by 'kubeconfig' connection")
}
