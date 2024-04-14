package console

import "fmt"

type Console struct{}

func (*Console) Log(msg ...interface{}) {
	fmt.Println(msg...)
}
