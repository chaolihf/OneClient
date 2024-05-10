package ui

//#cgo CFLAGS: -I../../cefapi
//#cgo LDFLAGS: -L${SRCDIR}/../../cefapi -L${SRCDIR}/../../Release -lcefapi -lcef
/*
#include <stdio.h>
#include "../../cefapi/cefapi.h"

*/
import "C"
import (
	"os"
	"unsafe"

	"github.com/lxn/win"
)

func TestMenuItem() {
	createBrowser("aaa", "http://www.sina.com.cn", 0, 0, 0, 0, 0)
}

func InitCef() {
	args := os.Args
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	C.setBeforePopupCallback(C.onBeforePopupFuncProto(C.cef_onBeforePopup))
	C.startCef(argc, (**C.char)(unsafe.Pointer(&argv[0])))
}

func ShutdownCef() {
	C.shutdownCef()
}

func createBrowser(title, url string, parent win.HWND, x, y, width, height int) {
	C.createBrowser(C.CString(title), C.CString(url), C.int(int(uintptr(parent))),
		C.int(x), C.int(y), C.int(width), C.int(height))
}

func loadUrl(url string) {
	C.loadUrl(C.CString(url))
}

func goBack() {
	C.goBack()
}
func goForward() {
	C.goForward()
}

func goReload() {
	C.goReload()
}

//export cef_onBeforePopup
func cef_onBeforePopup(target_url *C.char) C.int {
	width, height := window.GetSize()
	createBrowser("bbb", C.GoString(target_url), GetMainWindowHandler(), 0, toolbarHeight, width, height-toolbarHeight)
	return 1
}

func setBrowserSize(width, height int) {
	C.setBrowserSize(C.int(width), C.int(height))
}
