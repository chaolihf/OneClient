package plugin

import "go.uber.org/zap"

type CustomPlugins struct {
	logger  *zap.Logger
	plugins []IScriptPlugin
}

func NewCustomPlugins(logger *zap.Logger) *CustomPlugins {
	plugins := &CustomPlugins{
		logger: logger,
	}
	return plugins
}
