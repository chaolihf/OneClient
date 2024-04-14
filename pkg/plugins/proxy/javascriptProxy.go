package proxy

type JavascriptProxy struct{}

func (p *JavascriptProxy) CallPluginsMethod(code string, method string, parma interface{}) interface{} {
	return code
}
