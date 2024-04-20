package main

//https://github.com/mostafa/goja_debugger
import (
	"fmt"
	"log"
	"os"

	"com.chinatelecom.oneops.client/pkg/ui"
	"github.com/chaolihf/udpgo/lang"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = lang.InitProductLogger("logs/windows_helper.log", 300, 3, 10)
}

func main() {
	runner := NewJSRunner(logger)
	source := readData("asset/test.js")
	_, err := runner.runCode(source)
	if err != nil {
		panic(err)
	}
	v, err := runner.runFunction("ping", "134.64.116.90")
	fmt.Println(v.Export())
	v, err = runner.runFunction("robot")
	fmt.Println(v.Export())

	v, err = runner.runFunction("main", "http://134.64.116.90:8101/sso/index.html?res=workflow")
	fmt.Println(v.Export())
	v, err = runner.runFunction("main", "https://baidu.com")
	fmt.Println(v.Export())
	ui.ShowMain()
}

func readData(source string) string {
	data, err := os.ReadFile(source)
	if err != nil {
		log.Fatal(err)
	}
	// 将文件内容转换为字符串
	return string(data)
}
