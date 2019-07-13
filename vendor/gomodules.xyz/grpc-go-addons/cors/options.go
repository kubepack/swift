package cors

var (
	defaultOptions = &options{
		allowHost:      "*",
		allowSubdomain: true,
	}
)

type options struct {
	allowHost      string
	allowSubdomain bool
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type Option func(*options)

func OriginHost(host string) Option {
	return func(o *options) {
		o.allowHost = host
	}
}

func AllowSubdomain(allow bool) Option {
	return func(o *options) {
		o.allowSubdomain = allow
	}
}
