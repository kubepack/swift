package endpoints

type ProxyEndPoints []*endPoint

func (p *ProxyEndPoints) Register(fun interface{}) {
	if *p == nil {
		*p = make([]*endPoint, 0)
	}
	*p = append(*p, &endPoint{
		RegisterFunc: fun,
	})
}

var ProxyServerEndpoints = ProxyEndPoints{}
