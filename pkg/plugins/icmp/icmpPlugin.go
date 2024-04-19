/*
*
@description  实现IScriptPlugin接口的http插件
*/
package icmpPlugin

import (
	httpClient "com.chinatelecom.oneops.client/pkg/clients/httpclient"
)

// ICMPScriptPlugin 实现了IScriptPlugin接口的ICMP插件
type ICMPScriptPlugin struct {
}

// @return string icmp插件
func (hsp ICMPScriptPlugin) GetCode() string {
	return "icmp"
}

// CallPluginsMethod 实现IScriptPlugin接口的CallPluginsMethod方法
func (hsp ICMPScriptPlugin) CallPluginsMethod(method string, params interface{}) interface{} {
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
