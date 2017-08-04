package options

import (
	_env "github.com/appscode/go/env"
	"github.com/spf13/pflag"
)

type Config struct {
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
	MonitoringServerAddr string
}

func NewConfig() *Config {
	return &Config{
		SecureAddr:           ":50066",
		PlaintextAddr:        ":9866",
		// This Port Both Contains /metrics and /debug prefix with promethus and pprof.
		MonitoringServerAddr: ":6060",
	}
}

func (s *Config) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.SecureAddr, "secure-addr", s.SecureAddr, "host:port used to server secure apis")
	fs.StringVar(&s.PlaintextAddr, "plaintext-addr", s.PlaintextAddr, "host:port used to server plaintext apis")

	fs.BoolVar(&s.EnableJavaClient, "enable-java-client", s.EnableJavaClient, "Set true to send SETTINGS frame from the server. Default set to false")
	fs.StringVar(&s.APIDomain, "api-domain", s.APIDomain, "Domain used for apiserver (prod: api.appscode.com")
	fs.StringVar(&s.CACertFile, "caCertFile", s.CACertFile, "File containing CA certificate")
	fs.StringVar(&s.CertFile, "certFile", s.CertFile, "File container server TLS certificate")
	fs.StringVar(&s.KeyFile, "keyFile", s.KeyFile, "File containing server TLS private key")

	fs.BoolVar(&s.EnableCORS, "enable-cors", s.EnableCORS, "Enable CORS support")
	fs.StringVar(&s.CORSOriginHost, "cors-origin-host", s.CORSOriginHost, `Allowed CORS origin host e.g, domain[:port]`)
	fs.BoolVar(&s.CORSOriginAllowSubdomain, "cors-origin-allow-subdomain", s.CORSOriginAllowSubdomain, "Allow CORS request from subdomains of origin")

	fs.StringVar(&s.MonitoringServerAddr, "monitoring-addr", s.MonitoringServerAddr, "host:port used to serve as monitoring server")
}
