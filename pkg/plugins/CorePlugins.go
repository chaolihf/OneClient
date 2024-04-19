package plugin

import (
	httpPlugin "com.chinatelecom.oneops.client/pkg/plugins/http"
	icmpPlugin "com.chinatelecom.oneops.client/pkg/plugins/icmp"
	"go.uber.org/zap"
)

type CorePlugins struct {
	logger  *zap.Logger
	plugins []IScriptPlugin
}

func NewCorePlugins(logger *zap.Logger) *CorePlugins {
	plugins := &CorePlugins{
		logger: logger,
	}
	plugins.plugins = []IScriptPlugin{
		httpPlugin.NewHTTPScriptPlugin(logger),
		icmpPlugin.NewICMPScriptPlugin(logger),
	}
	return plugins
}
