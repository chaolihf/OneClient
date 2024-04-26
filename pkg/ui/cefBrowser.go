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

var isInit bool = false

func StartCefWindow(title string, url string) {
	args := os.Args
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	if !isInit {
		isInit = true
		C.startCef(argc, (**C.char)(unsafe.Pointer(&argv[0])))
	}
	if len(url) > 0 {
		C.createBrowser(C.CString(title), C.CString(url), 0)
	}
}

func createBrowser(title, url string, parent win.HWND) {
	C.createBrowser(C.CString(title), C.CString(url), C.int(int(uintptr(parent))))
}
