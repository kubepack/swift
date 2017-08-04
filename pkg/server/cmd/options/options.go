package options

import (
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
	WebAddr                  string
	EnableAnalytics          bool
}

func NewConfig() *Config {
	return &Config{
		SecureAddr:      ":50055",
		PlaintextAddr:   ":9855",
		WebAddr:         ":5050",
		EnableAnalytics: true,
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

	fs.StringVar(&s.WebAddr, "web-addr", s.WebAddr, "Address to listen on for web interface and telemetry.")
	fs.BoolVar(&s.EnableAnalytics, "analytics", s.EnableAnalytics, "Send analytical events to Google Analytics")
}
