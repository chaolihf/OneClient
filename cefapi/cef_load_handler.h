#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_load_handler_capi.h"
#include "utils.h"

///
  // Called when the loading state has changed. This callback will be executed
  // twice -- once when loading is initiated either programmatically or by user
  // action, and once when loading is terminated due to completion, cancellation
  // of failure. It will be called before any calls to OnLoadStart and after all
  // calls to OnLoadError and/or OnLoadEnd.
  ///
  void CEF_CALLBACK on_loading_state_change(struct _cef_load_handler_t* self,
                                              struct _cef_browser_t* browser,
                                              int isLoading,
                                              int canGoBack,
                                              int canGoForward){
        DEBUG_CALLBACK("on_loading_state_change\n");  
    }

///
  // Called after a navigation has been committed and before the browser begins
  // loading contents in the frame. The |frame| value will never be NULL -- call
  // the is_main() function to check if this frame is the main frame.
  // |transition_type| provides information about the source of the navigation
  // and an accurate value is only available in the browser process. Multiple
  // frames may be loading at the same time. Sub-frames may start or continue
  // loading after the main frame load has ended. This function will not be
  // called for same page navigations (fragments, history state, etc.) or for
  // navigations that fail or are canceled before commit. For notification of
  // overall browser load status use OnLoadingStateChange instead.
  ///
  void CEF_CALLBACK on_load_start(struct _cef_load_handler_t* self,
                                    struct _cef_browser_t* browser,
                                    struct _cef_frame_t* frame,
                                    cef_transition_type_t transition_type){
    if(frame){
        cef_string_userfree_t url=frame->get_url(frame);
        //char* url=ConvertCefStringToChar(url);
        DEBUG_CALLBACK("A");
    } 
    DEBUG_CALLBACK("on_load_start\n");  
  }

  ///
  // Called when the browser is done loading a frame. The |frame| value will
  // never be NULL -- call the is_main() function to check if this frame is the
  // main frame. Multiple frames may be loading at the same time. Sub-frames may
  // start or continue loading after the main frame load has ended. This
  // function will not be called for same page navigations (fragments, history
  // state, etc.) or for navigations that fail or are canceled before commit.
  // For notification of overall browser load status use OnLoadingStateChange
  // instead.
  ///
  void CEF_CALLBACK on_load_end(struct _cef_load_handler_t* self,
                                  struct _cef_browser_t* browser,
                                  struct _cef_frame_t* frame,
                                  int httpStatusCode){
    DEBUG_CALLBACK("on_load_end\n");  
  }

  ///
  // Called when a navigation fails or is canceled. This function may be called
  // by itself if before commit or in combination with OnLoadStart/OnLoadEnd if
  // after commit. |errorCode| is the error code number, |errorText| is the
  // error text and |failedUrl| is the URL that failed to load. See
  // net\base\net_error_list.h for complete descriptions of the error codes.
  ///
  void CEF_CALLBACK on_load_error(struct _cef_load_handler_t* self,
                                    struct _cef_browser_t* browser,
                                    struct _cef_frame_t* frame,
                                    cef_errorcode_t errorCode,
                                    const cef_string_t* errorText,
                                    const cef_string_t* failedUrl){
        DEBUG_CALLBACK("on_load_error\n");  
   }

void initialize_cef_load_handler(cef_load_handler_t *handler){
    DEBUG_CALLBACK("initialize_cef_load_handler\n");
    handler->base.size = sizeof(cef_load_handler_t);
    initialize_cef_base_ref_counted((cef_base_ref_counted_t*)handler);
    handler->on_loading_state_change= on_loading_state_change;
    handler->on_load_end=on_load_end;
    handler->on_load_error=on_load_error;
    handler->on_load_start=on_load_start;
}