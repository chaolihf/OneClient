package plugin

func GetAllPlugins() []IScriptPlugin {
	plugins := []IScriptPlugin{}
	plugins = append(plugins, NewCorePlugins().plugins...)
	plugins = append(plugins, NewCustomPlugins().plugins...)
	return plugins
}
