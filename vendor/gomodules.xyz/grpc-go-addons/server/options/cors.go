package options

import (
	"flag"

	"github.com/spf13/pflag"
	"gomodules.xyz/grpc-go-addons/server"
)

type CorsOptions struct {
	Enable         bool
	OriginHost     string
	AllowSubdomain bool
}

func NewCORSOptions() *CorsOptions {
	return &CorsOptions{
		OriginHost:     "*",
		AllowSubdomain: true,
	}
}

func (o *CorsOptions) AddGoFlags(fs *flag.FlagSet) {
	fs.BoolVar(&o.Enable, "enable-cors", o.Enable, "Enable CORS support")
	fs.StringVar(&o.OriginHost, "cors-origin-host", o.OriginHost, `Allowed CORS origin host e.g, domain[:port]`)
	fs.BoolVar(&o.AllowSubdomain, "cors-origin-allow-subdomain", o.AllowSubdomain, "Allow CORS request from subdomains of origin")
}

func (o *CorsOptions) AddFlags(fs *pflag.FlagSet) {
	gfs := flag.NewFlagSet("grpc-cors", flag.ExitOnError)
	o.AddGoFlags(gfs)
	fs.AddGoFlagSet(gfs)
}

func (o *CorsOptions) ApplyTo(cfg *server.Config) error {
	cfg.EnableCORS = o.Enable
	cfg.CORSOriginHost = o.OriginHost
	cfg.CORSAllowSubdomain = o.AllowSubdomain

	return nil
}

func (o *CorsOptions) Validate() []error {
	return nil
}
