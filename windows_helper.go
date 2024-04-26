package main

//https://github.com/mostafa/goja_debugger
import (
	"fmt"
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
	args := os.Args
	fmt.Println("start windows helper ", args)
	if len(args) == 1 {
		ui.ShowMain(logger)
	} else {
		ui.StartCefWindow("", "")
	}
}
