package plugin

import "go.uber.org/zap"

// IScriptPlugin 定义了一个包含CallMethod方法的接口，用于执行对服务的简单调用
type IScriptPlugin interface {

	// GetCode 返回插件代码
	GetCode() string

	/**
	 * 调用插件方法
	 * @param method 插件方法
	 * @param params 插件方法参数
	 * @return 返回插件方法执行结果
	 */
	CallPluginsMethod(method string, parmas interface{}) interface{}

	// SetLogger 设置日志
	SetLogger(logger *zap.Logger)
}
