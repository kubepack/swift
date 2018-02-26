package connectors

import "time"

type Config struct {
	Endpoint             string
	CACertFile           string
	ClientCertFile       string
	ClientPrivateKeyFile string
	InsecureSkipVerify   bool
	Timeout              time.Duration
	KubeContext          string
	LogRPC               bool
}
