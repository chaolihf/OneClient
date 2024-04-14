package main

//https://github.com/mostafa/goja_debugger
import (
	"fmt"

	"github.com/dop251/goja"
)

type Console struct{}

func (*Console) Log(msg ...interface{}) {
	fmt.Println(msg...)
}

func main() {
	vm := goja.New()
	vm.Set("console", &Console{})
	// 执行js字符串
	v, err := vm.RunString("console.Log('hello world js')")
	if err != nil {
		panic(err)
	}
	// export 获取执行结果
	fmt.Println(v.Export())

}
