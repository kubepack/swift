package endpoints

type GRPCEndpoints []*endPoint

func (s *GRPCEndpoints) Register(fun, server interface{}) {
	if *s == nil {
		*s = make([]*endPoint, 0)
	}
	*s = append(*s, &endPoint{
		RegisterFunc: fun,
		Server:       server,
	})
}

// all the public endpoints that will be exposed are listed
var GRPCServerEndpoints = GRPCEndpoints{}
