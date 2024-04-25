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
)

func StartCefWindow() {
	args := os.Args
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	C.startCef(argc, (**C.char)(unsafe.Pointer(&argv[0])))
}
