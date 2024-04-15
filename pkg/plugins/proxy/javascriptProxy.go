package proxy

import (
	httpClient "com.chinatelecom.oneops.client/pkg/clients/httpclient"
)

type JavascriptProxy struct{}

func (p *JavascriptProxy) CallPluginsMethod(code string, method string, parma interface{}) interface{} {
	client, err := httpClient.NewHttpClient()
	if err != nil {
		return nil
	} else {
		return client
	}
}
