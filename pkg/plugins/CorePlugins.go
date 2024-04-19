package plugin

import (
	httpPlugin "com.chinatelecom.oneops.client/pkg/plugins/http"
	icmpPlugin "com.chinatelecom.oneops.client/pkg/plugins/icmp"
)

type CorePlugins struct {
	plugins []IScriptPlugin
}

func NewCorePlugins() *CorePlugins {
	plugins := &CorePlugins{}
	plugins.plugins = []IScriptPlugin{
		httpPlugin.HTTPScriptPlugin{},
		icmpPlugin.ICMPScriptPlugin{},
	}
	return plugins
}
