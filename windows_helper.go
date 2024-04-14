package main

//https://github.com/mostafa/goja_debugger
import (
	"fmt"
	"log"
	"os"
)

func main() {
	runner := NewJSRunner()
	source := readData("asset/test.js")
	v, err := runner.runCode(source)
	if err != nil {
		panic(err)
	}
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
