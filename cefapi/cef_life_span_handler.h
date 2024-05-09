// CEF C API example
// Project website: https://github.com/cztomczak/cefcapi
//see https://zhuanlan.zhihu.com/p/619109619 for js binding
#pragma once

#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
#include "cefapi.h"

extern int g_browser_counter ;
extern cef_browser_t *g_browser;
extern onBeforePopupFuncProto onBeforePopupCallback;

// ----------------------------------------------------------------------------
// struct cef_life_span_handler_t
// ----------------------------------------------------------------------------

///
// Implement this structure to handle events related to browser life span. The
// functions of this structure will be called on the UI thread unless otherwise
// indicated.
///

// NOTE: There are many more callbacks in cef_life_span_handler,
//       but only on_before_close is implemented here.

///
// Called just before a browser is destroyed. Release all references to the
// browser object and do not attempt to execute any functions on the browser
// object after this callback returns. This callback will be the last
// notification that references |browser|. See do_close() documentation for
// additional usage information.
///
void CEF_CALLBACK on_before_close(struct _cef_life_span_handler_t* self,
                                  struct _cef_browser_t* browser) {
    DEBUG_CALLBACK("on_before_close\n");
    g_browser_counter--;
    if(g_browser_counter==0){
        cef_quit_message_loop();
    }
}

///
// Called after a new browser is created. It is now safe to begin performing
// actions with |browser|. cef_frame_handler_t callbacks related to initial
// main frame creation will arrive before this callback. See
// cef_frame_handler_t documentation for additional usage information.
///
void CEF_CALLBACK on_after_created (struct _cef_life_span_handler_t* self,
                                    struct _cef_browser_t* browser){
    DEBUG_CALLBACK("on_after_created\n");    
    if(g_browser== NULL){
        g_browser=browser;
        // cef_frame_t *mainFrame=g_browser->get_main_frame(g_browser);
        // cef_string_t url=getCefString("https://baidu.com");
        // mainFrame->load_url(mainFrame,&url);
    }
    g_browser_counter++;
}

///
  // Called on the UI thread before a new popup browser is created. The
  // |browser| and |frame| values represent the source of the popup request. The
  // |target_url| and |target_frame_name| values indicate where the popup
  // browser should navigate and may be NULL if not specified with the request.
  // The |target_disposition| value indicates where the user intended to open
  // the popup (e.g. current tab, new tab, etc). The |user_gesture| value will
  // be true (1) if the popup was opened via explicit user gesture (e.g.
  // clicking a link) or false (0) if the popup opened automatically (e.g. via
  // the DomContentLoaded event). The |popupFeatures| structure contains
  // additional information about the requested popup window. To allow creation
  // of the popup browser optionally modify |windowInfo|, |client|, |settings|
  // and |no_javascript_access| and return false (0). To cancel creation of the
  // popup browser return true (1). The |client| and |settings| values will
  // default to the source browser's values. If the |no_javascript_access| value
  // is set to false (0) the new browser will not be scriptable and may not be
  // hosted in the same renderer process as the source browser. Any
  // modifications to |windowInfo| will be ignored if the parent browser is
  // wrapped in a cef_browser_view_t. Popup browser creation will be canceled if
  // the parent browser is destroyed before the popup browser creation completes
  // (indicated by a call to OnAfterCreated for the popup browser). The
  // |extra_info| parameter provides an opportunity to specify extra information
  // specific to the created popup browser that will be passed to
  // cef_render_process_handler_t::on_browser_created() in the render process.
  ///
  int CEF_CALLBACK on_before_popup(
      struct _cef_life_span_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      const cef_string_t* target_url,
      const cef_string_t* target_frame_name,
      cef_window_open_disposition_t target_disposition,
      int user_gesture,
      const struct _cef_popup_features_t* popupFeatures,
      struct _cef_window_info_t* windowInfo,
      struct _cef_client_t** client,
      struct _cef_browser_settings_t* settings,
      struct _cef_dictionary_value_t** extra_info,
      int* no_javascript_access){
        DEBUG_CALLBACK("on_before_popup\n");  
        if (onBeforePopupCallback){
            char* url=convertCefStringToChar(target_url);
            int result=onBeforePopupCallback(url);
            free(url);
            return result;
        }
        return 0;
    }

life_span_handler_t * initialize_cef_life_span_handler() {
    DEBUG_CALLBACK("initialize_cef_life_span_handler\n");
    life_span_handler_t *h;
    h = calloc(1, sizeof(life_span_handler_t));
    cef_life_span_handler_t *handler = &h->handler;
    initialize_cef_base(h);
    handler->base.add_ref((cef_base_ref_counted_t *)h);
    handler->on_before_close = on_before_close;
    handler->on_after_created = on_after_created;
    handler->on_before_popup = on_before_popup;
    return h;
}
