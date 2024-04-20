package plugin

import (
	rpaPlugin "com.chinatelecom.oneops.client/pkg/plugins/rpa"
	"go.uber.org/zap"
)

type CustomPlugins struct {
	logger  *zap.Logger
	plugins []IScriptPlugin
}

func NewCustomPlugins(logger *zap.Logger) *CustomPlugins {
	plugins := &CustomPlugins{
		logger: logger,
	}
	plugins.plugins = []IScriptPlugin{
		rpaPlugin.NewRpaScriptPlugin(logger),
	}
	return plugins
}
