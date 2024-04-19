package proxy

import (
	"reflect"

	plugin "com.chinatelecom.oneops.client/pkg/plugins"
	"github.com/modern-go/reflect2"
)

type JavascriptProxy struct {
	plugins map[string]plugin.IScriptPlugin
}

func New() *JavascriptProxy {
	plugins := make(map[string]plugin.IScriptPlugin)
	for _, plugin := range plugin.GetAllPlugins() {
		plugins[plugin.GetCode()] = plugin
	}
	return &JavascriptProxy{
		plugins: plugins,
	}
}

/*
*
@param code 通过插件的包名和结构体名称来调用，如httpPlugin.HTTPScriptPlugin
*/
func (p *JavascriptProxy) CallPluginsByTypeName(code string, method string, parma interface{}) interface{} {
	pluginType := reflect2.TypeByName(code)
	obj := reflect.New(pluginType.Type1()).Elem().Interface()
	scriptPlugin, ok := obj.(plugin.IScriptPlugin)
	if ok {
		return scriptPlugin.CallPluginsMethod(method, parma)
	}
	return nil
}

func (p *JavascriptProxy) CallPluginsMethod(code string, method string, parma interface{}) interface{} {
	plugin, ok := p.plugins[code]
	if ok {
		return plugin.CallPluginsMethod(method, parma)
	}
	return nil
}
