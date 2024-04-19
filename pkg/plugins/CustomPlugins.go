package plugin

type CustomPlugins struct {
	plugins []IScriptPlugin
}

func NewCustomPlugins() *CustomPlugins {
	plugins := &CustomPlugins{}
	return plugins
}
