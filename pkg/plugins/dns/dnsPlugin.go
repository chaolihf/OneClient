/*
*
@description  实现IScriptPlugin接口的dns插件
*/
package dnsPlugin

import (
	"go.uber.org/zap"
)

// DnsScriptPlugin 实现了IScriptPlugin接口的dns插件
type DnsScriptPlugin struct {
	logger *zap.Logger
}

func NewDnsScriptPlugin(logger *zap.Logger) *DnsScriptPlugin {
	return &DnsScriptPlugin{logger: logger}
}

// @return string dns插件
func (thisPlugin DnsScriptPlugin) GetCode() string {
	return "dns"
}

// CallPluginsMethod 实现IScriptPlugin接口的CallPluginsMethod方法
func (thisPlugin DnsScriptPlugin) CallPluginsMethod(method string, params interface{}) interface{} {
	switch method {
	case "probe":
		{
			thisPlugin.probe(params.(string))
		}
	}
	return nil
}

func (thisPlugin *DnsScriptPlugin) SetLogger(logger *zap.Logger) {
	thisPlugin.logger = logger
}

func (thisPlugin *DnsScriptPlugin) probe(target string) bool {
	return false
}
