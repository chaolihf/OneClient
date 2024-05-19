#pragma once

typedef signed char Go_Int8;
typedef unsigned char Go_Uint8;
typedef short Go_Int16;
typedef unsigned short Go_Uint16;
typedef int Go_Int32;
typedef unsigned int Go_Uint32;
typedef long long Go_Int64;
typedef unsigned long long Go_Uint64;
typedef Go_Int64 Go_Int;
typedef Go_Uint64 Go_Uint;
typedef size_t Go_Uintptr;
typedef float Go_Float32;
typedef double Go_Float64;
typedef struct { const char *p; ptrdiff_t n; } Go_String;

#ifndef GO_CGO_EXPORT_PROLOGUE_H
struct cef_onResourceHandlerGetResponseHeaders_return {
	Go_Int status;
	Go_String mime_type;
    Go_Int response_length;
};

struct cef_onResourceHandlerRead_return {
	Go_Int bytes_read;
	Go_Int has_data;
} ;
#endif

void goCopyMemory(void* _Dst,void const* _Src,int _Size);

int startCef(int argc, char** argv) ;

int number_add_mod(int a, int b, int mod);

void shutdownCef();

int createBrowser(const char * title,const char * url,int parent_window_handle,int x,int y,int width,int height);

void loadUrl(const char * url);

void goBack();

void goForward();

void goReload();

void setForegroundWindow(int window_handle);

void setBrowserSize(int width, int height);

#ifndef __TEST_H__
#define __TEST_H__
#ifdef __cplusplus
extern "C"{
#endif

//define function pointer
typedef int(*onBeforePopupFuncProto) (char *target_url);
typedef int(*onResourceHandlerOpenFuncProto) (char *target_url,int request_id);
typedef struct cef_onResourceHandlerGetResponseHeaders_return(*onResourceHandlerGetResponseHeadersFuncProto) (int request_id);
typedef struct cef_onResourceHandlerRead_return(*onResourceHandlerReadFuncProto) (int request_id,void* data_out,int bytes_to_read);

//setup callback function
void setBeforePopupCallback(onBeforePopupFuncProto s);
void setResourceHandlerOpenCallback(onResourceHandlerOpenFuncProto s);
void setResourceHandlerGetResponseHeadersCallback(onResourceHandlerGetResponseHeadersFuncProto s);
void setResourceHandlerReadCallback(onResourceHandlerReadFuncProto s);


//define in go function
#ifndef GO_CGO_EXPORT_PROLOGUE_H
int cef_onBeforePopup(char *);
int cef_onResourceHandlerOpen(char *,int);
struct cef_onResourceHandlerGetResponseHeaders_return cef_onResourceHandlerGetResponseHeaders(int);
struct cef_onResourceHandlerRead_return cef_onResourceHandlerRead(int,void*,int);
#endif


#ifdef __cplusplus
}
#endif
#endif