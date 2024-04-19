package plugin

import (
	httpPlugin "com.chinatelecom.oneops.client/pkg/plugins/http"
)

type CorePlugins struct {
	plugins []IScriptPlugin
}

func NewCorePlugins() *CorePlugins {
	plugins := &CorePlugins{}
	plugins.plugins = []IScriptPlugin{
		httpPlugin.HTTPScriptPlugin{},
	}
	return plugins
}
