package main

//#cgo CFLAGS: -I./cefapi
//#cgo LDFLAGS: -L${SRCDIR}/cefapi -L${SRCDIR}/Release -lcefapi -lcef
/*
#include <stdio.h>
#include "cefapi.h"
*/
import "C"

//https://github.com/mostafa/goja_debugger
import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"com.chinatelecom.oneops.client/pkg/ui"
	"github.com/chaolihf/udpgo/lang"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = lang.InitProductLogger("logs/windows_helper.log", 300, 3, 10)
}

func main() {
	fmt.Println("start windows helper")
	args := os.Args
	if len(args) == 1 {
		go runTest()
	}
	if len(args) >= 1 {
		argc := C.int(len(args))
		argv := make([]*C.char, argc)
		for i, arg := range args {
			argv[i] = C.CString(arg)
		}
		C.startCef(argc, (**C.char)(unsafe.Pointer(&argv[0])))
	}
}

func runTest() {
	ui.ShowMain()
	runner := NewJSRunner(logger)
	source := readData("../asset/test.js")
	_, err := runner.runCode(source)
	if err != nil {
		panic(err)
	}
	fmt.Println("start run test")
	v, err := runner.runFunction("ping", "134.64.116.90")
	fmt.Println(v.Export())
	v, err = runner.runFunction("robot")
	fmt.Println(v.Export())

	v, err = runner.runFunction("main", "http://134.64.116.90:8101/sso/index.html?res=workflow")
	fmt.Println(v.Export())
	v, err = runner.runFunction("main", "https://baidu.com")
	fmt.Println(v.Export())
}

func readData(source string) string {
	data, err := os.ReadFile(source)
	if err != nil {
		log.Fatal(err)
	}
	// 将文件内容转换为字符串
	return string(data)
}
