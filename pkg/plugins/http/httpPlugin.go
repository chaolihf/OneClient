/*
*
@description  实现IScriptPlugin接口的http插件
*/
package httpPlugin

import (
	httpClient "com.chinatelecom.oneops.client/pkg/clients/httpclient"
	"go.uber.org/zap"
)

// HTTPScriptPlugin 实现了IScriptPlugin接口的HTTP插件
type HTTPScriptPlugin struct {
	logger *zap.Logger
}

func NewHTTPScriptPlugin(logger *zap.Logger) *HTTPScriptPlugin {
	return &HTTPScriptPlugin{logger: logger}
}

// @return string http插件
func (thisPlugin HTTPScriptPlugin) GetCode() string {
	return "http"
}

// CallPluginsMethod 实现IScriptPlugin接口的CallPluginsMethod方法
func (thisPlugin HTTPScriptPlugin) CallPluginsMethod(method string, params interface{}) interface{} {
	switch method {
	case "newClient":
		{
			client, err := httpClient.NewHttpClient()
			if err != nil {
				return nil
			} else {
				return client
			}
		}
	}
	return nil
}

func (thisPlugin *HTTPScriptPlugin) SetLogger(logger *zap.Logger) {
	thisPlugin.logger = logger
}
