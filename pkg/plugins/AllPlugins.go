package plugin

import "go.uber.org/zap"

func GetAllPlugins(logger *zap.Logger) []IScriptPlugin {
	plugins := []IScriptPlugin{}
	plugins = append(plugins, NewCorePlugins(logger).plugins...)
	plugins = append(plugins, NewCustomPlugins(logger).plugins...)
	return plugins
}
