#include <stddef.h>
int main() { return 0; }
void crosscall2(void(*fn)(void*) __attribute__((unused)), void *a __attribute__((unused)), int c __attribute__((unused)), size_t ctxt __attribute__((unused))) { }
size_t _cgo_wait_runtime_init_done(void) { return 0; }
void _cgo_release_context(size_t ctxt __attribute__((unused))) { }
char* _cgo_topofstack(void) { return (char*)0; }
void _cgo_allocate(void *a __attribute__((unused)), int c __attribute__((unused))) { }
void _cgo_panic(void *a __attribute__((unused)), int c __attribute__((unused))) { }
void _cgo_reginit(void) { }
#line 1 "cgo-generated-wrappers"
extern void cef_onBeforePopup();
extern void cef_onResourceHandlerOpen();
void _cgoexp_8c88dd3c38f3_cef_onBeforePopup(void* p){}
void _cgoexp_8c88dd3c38f3_cef_onResourceHandlerOpen(void* p){}
void _cgoexp_8c88dd3c38f3_CopyDataToMemory(void* p){}
void _cgoexp_8c88dd3c38f3_cef_onResourceHandlerGetResponseHeaders(void* p){}
void _cgoexp_8c88dd3c38f3_cef_onResourceHandlerRead(void* p){}
