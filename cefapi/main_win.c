// CEF C API example
// Project website: https://github.com/cztomczak/cefcapi

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdbool.h>
#include "cefapi/cef_base.h"
#include "cefapi/cef_app.h"
#include "cefapi/cef_client.h"
#include "cefapi/cef_life_span_handler.h"
#include "cefapi/cef_load_handler.h"
#include "cefapi/cef_render_process_handler.h"
#include "cefapi/cef_display_handler.h"

#include "include/cef_version.h"

#include "cefapi.h"
#include "utils.h"
#include <windows.h>


cef_browser_settings_t g_browser_settings = {};
bool isStartMessageLoop=false;

int globalValue=0;
int number_add_mod(int a, int b, int mod) {
    globalValue++;
    return (a+b)%mod+globalValue;
}

// Globals
cef_browser_t *g_browser;
client_t *g_client;
life_span_handler_t *g_life_span_handler;
load_handler *g_load_handler;
//display_handler *g_display_handler;
cef_display_handler_t g_display_handler={};

cef_render_process_handler_t g_cef_render_process_handler={};
int g_browser_counter=0;
bool isBrowserProcess=false;
#ifdef windowsapp
int main(int argc, char** argv) {
    return startCef(argc, argv);
}
#endif

//all callback definition
onBeforePopupFuncProto onBeforePopupCallback;


int startCef(int argc, char** argv) {
    // This executable is called many times, because it
    // is also used for subprocesses. Let's print args
    // so we can differentiate between main process and
    // subprocesses. If one of the first args is for
    // example "--type=renderer" then it means that
    // this is a Renderer process. There may be more
    // subprocesses like GPU (--type=gpu-process) and
    // others. On Linux there are also special Zygote
    // processes.
    printf("\nProcess args: ");
    if (argc == 1) {
        printf("none (Main process)");
        isBrowserProcess=true;
    } else {
        for (int i = 1; i < argc; i++) {
            if (strlen(argv[i]) > 128){
                printf("... ");
            }
            else{
                printf("%s ", argv[i]);
                if (strcmp(argv[i], "--type=renderer") == 0) {
                    DWORD processId = GetCurrentProcessId();
                    char message[256];
                    sprintf(message, "render process id : %lu", processId);
                    //for debug
                    //MessageBox(NULL, message, "debug process", MB_OK);
                    break;
                }
            }
        }
    }
    printf("\n\n");

    // CEF version
    if (argc == 1) {
        printf("CEF version: %s\n", CEF_VERSION);
    }

    // Main args
    cef_main_args_t main_args = {};
    main_args.instance = GetModuleHandle(NULL);

    //g_cef_render_process_handler=initialize_cef_render_process_handler();
    initialize_cef_render_process_handler_direct(&g_cef_render_process_handler);

    // Cef app
    app_t *app=initialize_cef_app();
    
    // Execute subprocesses. It is also possible to have
    // a separate executable for subprocesses by setting
    // cef_settings_t.browser_subprocess_path. In such
    // case cef_execute_process should not be called here.
    printf("cef_execute_process, argc=%d\n", argc);
    int code = cef_execute_process(&main_args, (cef_app_t *)app, NULL);
    if (code >= 0) {
        _exit(code);
    }
    int result;
    // Application settings. It is mandatory to set the
    // "size" member.
    cef_settings_t settings = {};
    //不知道为什么在3版本中直接计算sizeof大小，但在最新版本中需要减去16来计算大小；
    //don't known why should substrace 16 byte for size,otherwise will cause invalid settings->[base.]size
    settings.size = sizeof(cef_settings_t);
    settings.log_severity = LOGSEVERITY_ERROR; // Show only warnings/errors
    settings.no_sandbox = 1;

    // Initialize CEF
    printf("cef_initialize\n");
    app=initialize_cef_app();
    cef_app_t *cef_app = (cef_app_t *)app;
    cef_app->base.add_ref((cef_base_ref_counted_t *)app);
    result=cef_initialize(&main_args, &settings, cef_app, NULL);
    cef_app->base.release((cef_base_ref_counted_t *)app);
    if(result==0){
        printf("cef_initialize failed\n");
        return 0;
    }
    // Browser settings. It is mandatory to set the
    // "size" member.
    g_browser_settings.size = sizeof(cef_browser_settings_t);
    
    // Client handlers
    g_client=initialize_cef_client();
    g_life_span_handler=initialize_cef_life_span_handler();
    g_load_handler=initialize_cef_load_handler();
    //g_display_handler=initialize_display_handler();
    initialize_display_handler_direct(&g_display_handler);
    #ifdef windowsapp
        createBrowser("aa","http://baidu.com",0,0,0,0,0);
    #endif  
    //
    if (!isStartMessageLoop){
        isStartMessageLoop=true;
        // Message loop. There is also cef_do_message_loop_work()
        // that allow for integrating with existing message loops.
        // On Windows for best performance you should set
        // cef_settings_t.multi_threaded_message_loop to true.
        // Note however that when you do that CEF UI thread is no
        // more application main thread and using CEF API is more
        // difficult and require using functions like cef_post_task
        // for running tasks on CEF UI thread.
        printf("cef_run_message_loop\n");
        cef_run_message_loop();
    }
    return 0;
}


void shutdownCef(){
    // Shutdown CEF
    printf("cef_shutdown\n");
    cef_shutdown();
}

int createBrowser(const char * title,const char * url,int parent_window_handle,int x,int y,int width,int height){
    // Window info
    cef_window_info_t window_info = {};
    window_info.style =  WS_CLIPCHILDREN \
            | WS_CLIPSIBLINGS | WS_VISIBLE  ;
    if(parent_window_handle!=0){
        window_info.style=window_info.style | WS_CHILD |
        WS_OVERLAPPED | WS_THICKFRAME;
        window_info.bounds.x = x;
        window_info.bounds.y = y;
        window_info.bounds.width = width;
        window_info.bounds.height = height;
    }
    else {
        window_info.style=window_info.style | WS_OVERLAPPEDWINDOW ;
        window_info.bounds.x = CW_USEDEFAULT;
        window_info.bounds.y = CW_USEDEFAULT;
        window_info.bounds.width = CW_USEDEFAULT;
        window_info.bounds.height = CW_USEDEFAULT;
    }
    window_info.ex_style=WS_EX_TOPMOST;    

    // Window info - window title
    window_info.window_name = getCefString(title);
    window_info.parent_window=(HWND)(intptr_t)parent_window_handle;
    // Initial url
    cef_string_t cef_url = getCefString(url);

    // Create browser asynchronously. There is also a
    // synchronous version of this function available.
    printf("cef_browser_host_create_browser\n");
    int result=cef_browser_host_create_browser(&window_info, (cef_client_t *)g_client, &cef_url,
                                    &g_browser_settings, NULL,NULL);
    if(result==0)
    {
        printf("cef_browser_host_create_browser failed\n");
    }
    // cef_browser_t *result=cef_browser_host_create_browser_sync(&window_info, &g_client, &cef_url,
    //                                 &g_browser_settings, NULL,NULL);
    // return 0;
    return result;
    
}

void loadUrl(const char* url){
    if (isBrowserProcess){
        cef_frame_t *frame=g_browser->get_main_frame(g_browser);
        cef_string_userfree_t oldUrl=frame->get_url(frame);
        cef_string_utf8_t cefString={};
        cef_string_utf16_to_utf8(oldUrl->str,oldUrl->length,&cefString);
        char* result=(char *)calloc(1,cefString.length+1);
        memcpy(result,cefString.str,cefString.length);
        result[cefString.length]='\0';
        printf("oldUrl:%s\n",result);
        free(result);


        const cef_string_t cef_url = getCefString(url);
        // cef_process_message_t *message=cef_process_message_create(&cef_url);
        // cef_list_value_t *arguments= message->get_argument_list(message);
        // arguments->set_int(arguments,0,10);
        // frame->send_process_message(frame,PID_RENDERER,message);
        //frame->select_all(frame);
        frame->load_url(frame,&cef_url);
        
    }
}

void goBack(){
    if (isBrowserProcess){
        g_browser->go_back(g_browser);
    }
}

void goForward(){
    if (isBrowserProcess){
        g_browser->go_forward(g_browser);
    }
}

void goReload(){
    if (isBrowserProcess){
        g_browser->reload(g_browser);
    }
}

void setForegroundWindow(int window_handle){
    SetForegroundWindow((HWND)(intptr_t)window_handle);
}


void setBeforePopupCallback(onBeforePopupFuncProto callback){
    onBeforePopupCallback=callback;
}
