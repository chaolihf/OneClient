/*
*
@description  实现IScriptPlugin接口的http插件
*/
package httpPlugin

import (
	httpClient "com.chinatelecom.oneops.client/pkg/clients/httpclient"
)

// HTTPScriptPlugin 实现了IScriptPlugin接口的HTTP插件
type HTTPScriptPlugin struct {
	// 在这里添加任何必要的成员变量，例如API URL、客户端配置等
}

// @return string http插件
func (hsp HTTPScriptPlugin) GetCode() string {
	return "http"
}

// CallPluginsMethod 实现IScriptPlugin接口的CallPluginsMethod方法
func (hsp HTTPScriptPlugin) CallPluginsMethod(method string, params interface{}) interface{} {
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
