package ui

//#cgo CFLAGS: -I../../cefapi
//#cgo LDFLAGS: -L${SRCDIR}/../../cefapi -L${SRCDIR}/../../Release -lcefapi -lcef
/*
#include <stdio.h>
#include <stdlib.h>

#include "../../cefapi/cefapi.h"

*/
import "C"
import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"github.com/lxn/win"
)

var mainPageContent = "<html><head><title>Scheme Test(Home Page) From Go</title></head><body bgcolor=\"white\">This contents of this page page are served by the resource handlle (golang version), navigate to new page <a href='https://www.sina.com.cn'>Sina</></body></html>"
var offset_ int = 0

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
	C.setResourceHandlerOpenCallback(C.onResourceHandlerOpenFuncProto(C.cef_onResourceHandlerOpen))
	C.setResourceHandlerGetResponseHeadersCallback(C.onResourceHandlerGetResponseHeadersFuncProto(C.cef_onResourceHandlerGetResponseHeaders))
	C.setResourceHandlerReadCallback(C.onResourceHandlerReadFuncProto(C.cef_onResourceHandlerRead))
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

//export cef_onResourceHandlerOpen
func cef_onResourceHandlerOpen(target_url *C.char, identity C.int) C.int {
	fmt.Printf("open resource %s with identity %d\n", C.GoString(target_url), identity)
	return 1
}

func setBrowserSize(width, height int) {
	C.setBrowserSize(C.int(width), C.int(height))
}

//export cef_onResourceHandlerGetResponseHeaders
func cef_onResourceHandlerGetResponseHeaders(identity C.int) (int, string, int) {
	return 200, "text/html", len([]byte(mainPageContent))
}

//export cef_onResourceHandlerRead
func cef_onResourceHandlerRead(identity C.int, data_out *C.char, bytes_to_read C.int) (int, int) {
	bytes_read := 0
	has_data := 0
	sourceData := []byte(mainPageContent)
	if offset_ < len(sourceData) {
		transfer_size := int(math.Min(float64(bytes_to_read), float64(len(sourceData)-offset_)))
		C.goCopyMemory(unsafe.Pointer(data_out), unsafe.Pointer(&sourceData[offset_]), C.int(transfer_size))
		offset_ = offset_ + transfer_size
		bytes_read = transfer_size
		has_data = 1
	}
	return bytes_read, has_data
}
