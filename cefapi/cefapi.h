#pragma once

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

//setup callback function
void setBeforePopupCallback(onBeforePopupFuncProto s);
void setResourceHandlerOpenCallback(onResourceHandlerOpenFuncProto s);

//define in go function
int cef_onBeforePopup(char *);
int cef_onResourceHandlerOpen(char *,int);


#ifdef __cplusplus
}
#endif
#endif