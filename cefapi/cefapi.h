#pragma once

#ifndef GO_CGO_EXPORT_PROLOGUE_H
typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef size_t GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;

//#ifndef GO_CGO_GOSTRING_TYPEDEF
	typedef struct { const char *p; ptrdiff_t n; } Go_String;
//#endif

struct cef_onResourceHandlerGetResponseHeaders_return {
	GoInt status;
	Go_String mime_type;
    GoInt response_length;
};

struct cef_onResourceHandlerRead_return {
	GoInt bytes_read;
	GoInt has_data;
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

void goSendMouseEvent();

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