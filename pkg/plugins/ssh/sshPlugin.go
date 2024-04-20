/*
*
@description  实现IScriptPlugin接口的ssh插件
*/
package sshPlugin

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

// SshScriptPlugin 实现了IScriptPlugin接口的ssh插件
type SshScriptPlugin struct {
	logger *zap.Logger
}

func NewSshScriptPlugin(logger *zap.Logger) *SshScriptPlugin {
	return &SshScriptPlugin{logger: logger}
}

// @return string ssh插件
func (thisPlugin SshScriptPlugin) GetCode() string {
	return "ssh"
}

// CallPluginsMethod 实现IScriptPlugin接口的CallPluginsMethod方法
func (thisPlugin SshScriptPlugin) CallPluginsMethod(method string, params interface{}) interface{} {
	switch method {
	case "newClient":
		{
			thisPlugin.probe(params.(string))
		}
	}
	return nil
}

func (thisPlugin *SshScriptPlugin) SetLogger(logger *zap.Logger) {
	thisPlugin.logger = logger
}

func (thisPlugin *SshScriptPlugin) probe(target string) bool {
	client, err := ssh.Dial("tcp", target, config)
	return false
}
