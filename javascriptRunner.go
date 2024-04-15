package main

import (
	"com.chinatelecom.oneops.client/pkg/plugins/console"
	"com.chinatelecom.oneops.client/pkg/plugins/proxy"
	"github.com/dop251/goja"
)

type JSRunner struct {
	runtime *goja.Runtime
}

func NewJSRunner() *JSRunner {
	vm := goja.New()
	vm.Set("console", &console.Console{})
	vm.Set("_func_predef_proxy", &proxy.JavascriptProxy{})
	return &JSRunner{
		runtime: vm,
	}
}

func (runner *JSRunner) runCode(str string) (goja.Value, error) {
	return runner.runtime.RunString(str)
}

func (runner *JSRunner) runFunction(name string, args ...interface{}) (goja.Value, error) {
	function, ok := goja.AssertFunction(runner.runtime.Get(name))
	var values []goja.Value
	for _, v := range args {
		values = append(values, runner.runtime.ToValue(v))
	}
	if ok {
		return function(goja.Undefined(), values...)
	} else {
		return nil, &goja.InterruptedError{}
	}
}
