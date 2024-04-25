// CEF C API example
// Project website: https://github.com/cztomczak/cefcapi

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>

#include "cefapi/cef_base.h"
#include "cefapi/cef_app.h"
#include "cefapi/cef_client.h"
#include "cefapi/cef_life_span_handler.h"
#include "include/cef_version.h"

#include "cefapi.h"
#include "utils.h"

int number_add_mod(int a, int b, int mod) {
    return (a+b)%mod;
}

// Globals
cef_life_span_handler_t g_life_span_handler = {};
#ifdef windowsapp
int main(int argc, char** argv) {
    return startCef(argc, argv);
}
#endif
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
    } else {
        for (int i = 1; i < argc; i++) {
            if (strlen(argv[i]) > 128)
                printf("... ");
            else
                printf("%s ", argv[i]);
        }
    }
    printf("\n\n");

    // CEF version
    if (argc == 0) {
        printf("CEF version: %s\n", CEF_VERSION);
    }

    // Main args
    cef_main_args_t main_args = {};
    main_args.instance = GetModuleHandle(NULL);

    // Cef app
    cef_app_t app = {};
    initialize_cef_app(&app);
    
    // Execute subprocesses. It is also possible to have
    // a separate executable for subprocesses by setting
    // cef_settings_t.browser_subprocess_path. In such
    // case cef_execute_process should not be called here.
    printf("cef_execute_process, argc=%d\n", argc);
    int code = cef_execute_process(&main_args, &app, NULL);
    if (code >= 0) {
        _exit(code);
    }
    int result;
    // Application settings. It is mandatory to set the
    // "size" member.
    cef_settings_t settings = {};
    //不知道为什么在3版本中直接计算sizeof大小，但在最新版本中需要减去16来计算大小；
    //don't known why should substrace 16 byte for size,otherwise will cause invalid settings->[base.]size
    settings.size = sizeof(cef_settings_t)-16;
    settings.log_severity = LOGSEVERITY_VERBOSE; // Show only warnings/errors
    settings.no_sandbox = 1;

    // Initialize CEF
    printf("cef_initialize\n");
    result=cef_initialize(&main_args, &settings, &app, NULL);
    if(result==0){
        printf("cef_initialize failed\n");
        return 0;
    }
    // Window info
    cef_window_info_t window_info = {};
    window_info.style = WS_OVERLAPPEDWINDOW | WS_CLIPCHILDREN \
            | WS_CLIPSIBLINGS | WS_VISIBLE;
    window_info.parent_window = NULL;
    window_info.bounds.x = CW_USEDEFAULT;
    window_info.bounds.y = CW_USEDEFAULT;
    window_info.bounds.width = CW_USEDEFAULT;
    window_info.bounds.height = CW_USEDEFAULT;

    // Window info - window title
    window_info.window_name = getCefString("cefcapi example");

    // Initial url
    cef_string_t cef_url = getCefString("http://baidu.com");

    // Browser settings. It is mandatory to set the
    // "size" member.
    cef_browser_settings_t browser_settings = {};
    browser_settings.size = sizeof(cef_browser_settings_t)-16;
    
    // Client handlers
    cef_client_t client = {};
    initialize_cef_client(&client);
    initialize_cef_life_span_handler(&g_life_span_handler);

    // Create browser asynchronously. There is also a
    // synchronous version of this function available.
    printf("cef_browser_host_create_browser\n");
    result=cef_browser_host_create_browser(&window_info, &client, &cef_url,
                                    &browser_settings, NULL,NULL);
    if(result==0)
    {
        printf("cef_browser_host_create_browser failed\n");
        return 0;
    }
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

    // Shutdown CEF
    printf("cef_shutdown\n");
    cef_shutdown();

    return 0;
}

