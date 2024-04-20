/*
*
@description  实现IScriptPlugin接口的tcp插件
*/
package tcpPlugin

import (
	"go.uber.org/zap"
)

// TcpScriptPlugin 实现了IScriptPlugin接口的tcp插件
type TcpScriptPlugin struct {
	logger *zap.Logger
}

func NewTcpScriptPlugin(logger *zap.Logger) *TcpScriptPlugin {
	return &TcpScriptPlugin{logger: logger}
}

// @return string tcp插件
func (thisPlugin TcpScriptPlugin) GetCode() string {
	return "tcp"
}

// CallPluginsMethod 实现IScriptPlugin接口的CallPluginsMethod方法
func (thisPlugin TcpScriptPlugin) CallPluginsMethod(method string, params interface{}) interface{} {
	switch method {
	case "probe":
		{
			thisPlugin.probe(params.(string))
		}
	}
	return nil
}

func (thisPlugin *TcpScriptPlugin) SetLogger(logger *zap.Logger) {
	thisPlugin.logger = logger
}

func (thisPlugin *TcpScriptPlugin) probe(target string) bool {
	return false
}
