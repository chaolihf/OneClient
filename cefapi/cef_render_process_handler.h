#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_load_handler_capi.h"
#include "utils.h"


  ///
  // Called after a browser has been created. When browsing cross-origin a new
  // browser will be created before the old browser with the same identifier is
  // destroyed. |extra_info| is an optional read-only value originating from
  // cef_browser_host_t::cef_browser_host_create_browser(),
  // cef_browser_host_t::cef_browser_host_create_browser_sync(),
  // cef_life_span_handler_t::on_before_popup() or
  // cef_browser_view_t::cef_browser_view_create().
  ///
  void CEF_CALLBACK on_browser_created(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_dictionary_value_t* extra_info){
    DEBUG_CALLBACK("on_browser_created----------------\n");

  }
void initialize_cef_render_process_handler(cef_render_process_handler_t *handler){
    DEBUG_CALLBACK("initialize_cef_render_process_handler\n");
    handler->base.size = sizeof(cef_render_process_handler_t);
    initialize_cef_base_ref_counted((cef_base_ref_counted_t*)handler);
    handler->on_browser_created= on_browser_created;
}