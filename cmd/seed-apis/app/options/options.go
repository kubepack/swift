package options

import (
	_env "github.com/appscode/go/env"
	"github.com/spf13/pflag"
)

type Config struct {
	APIPort          int
	PprofPort        int
	ReportMonitoring bool
	TokenSecret      string
}

func NewConfig() *Config {
	return &Config{
		APIPort:          50066,
		PprofPort:        6060,
		ReportMonitoring: !_env.FromHost().DevMode(),
	}
}

func (s *Config) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar(&s.APIPort, "api-port", s.APIPort, "Port used to serve apis")
	fs.IntVar(&s.PprofPort, "pprof-port", s.PprofPort, "port used to run pprof tools")
	fs.BoolVar(&s.ReportMonitoring, "report-monitoring", s.ReportMonitoring, "Report monitoring, disabled for dev env by default")
}
