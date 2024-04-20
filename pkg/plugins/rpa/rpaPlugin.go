/*
@description  实现IScriptPlugin接口的rpa插件
*/
package rpaPlugin

import (
	"github.com/go-vgo/robotgo"
	"go.uber.org/zap"
)

// RpaScriptPlugin 实现了IScriptPlugin接口的rpa插件
type RpaScriptPlugin struct {
	logger *zap.Logger
}

func NewRpaScriptPlugin(logger *zap.Logger) *RpaScriptPlugin {
	return &RpaScriptPlugin{logger: logger}
}

// @return string rpa插件
func (thisPlugin RpaScriptPlugin) GetCode() string {
	return "rpa"
}

// CallPluginsMethod 实现IScriptPlugin接口的CallPluginsMethod方法
func (thisPlugin RpaScriptPlugin) CallPluginsMethod(method string, params interface{}) interface{} {
	switch method {
	case "probe":
		{
			thisPlugin.probe()
		}
	}
	return nil
}

func (thisPlugin *RpaScriptPlugin) SetLogger(logger *zap.Logger) {
	thisPlugin.logger = logger
}

func (thisPlugin *RpaScriptPlugin) probe() bool {
	robotgo.MouseSleep = 100

	robotgo.ScrollDir(10, "up")
	robotgo.ScrollDir(20, "right")

	robotgo.Scroll(0, -10)
	robotgo.Scroll(100, 0)

	robotgo.MilliSleep(100)
	robotgo.ScrollSmooth(-10, 6)
	// robotgo.ScrollRelative(10, -100)

	robotgo.Move(10, 20)
	robotgo.MoveRelative(0, -10)
	robotgo.DragSmooth(10, 10)

	robotgo.Click("wheelRight")
	robotgo.Click("left", true)
	robotgo.MoveSmooth(100, 200, 1.0, 10.0)

	robotgo.Toggle("left")
	robotgo.Toggle("left", "up")
	return false
}
