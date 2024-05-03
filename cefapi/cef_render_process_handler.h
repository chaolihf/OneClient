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


  ///
  // Called before a browser is destroyed.
  ///
  void CEF_CALLBACK on_browser_destroyed(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser){
        DEBUG_CALLBACK("on_browser_destroyed\n");
      }


  ///
  // Called immediately after the V8 context for a frame has been created. To
  // retrieve the JavaScript 'window' object use the
  // cef_v8context_t::get_global() function. V8 handles can only be accessed
  // from the thread on which they are created. A task runner for posting tasks
  // on the associated thread can be retrieved via the
  // cef_v8context_t::get_task_runner() function.
  ///
  void CEF_CALLBACK on_context_created(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      struct _cef_v8context_t* context){
        DEBUG_CALLBACK("on_context_created\n");
      }

  ///
  // Called immediately before the V8 context for a frame is released. No
  // references to the context should be kept after this function is called.
  ///
  void CEF_CALLBACK on_context_released(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      struct _cef_v8context_t* context){
        DEBUG_CALLBACK("on_context_released\n");
      }

  ///
  // Called for global uncaught exceptions in a frame. Execution of this
  // callback is disabled by default. To enable set
  // CefSettings.uncaught_exception_stack_size > 0.
  ///
  void CEF_CALLBACK on_uncaught_exception(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      struct _cef_v8context_t* context,
      struct _cef_v8exception_t* exception,
      struct _cef_v8stack_trace_t* stackTrace){
        DEBUG_CALLBACK("on_uncaught_exception\n");
      }

  ///
  // Called when a new node in the the browser gets focus. The |node| value may
  // be NULL if no specific node has gained focus. The node object passed to
  // this function represents a snapshot of the DOM at the time this function is
  // executed. DOM objects are only valid for the scope of this function. Do not
  // keep references to or attempt to access any DOM objects outside the scope
  // of this function.
  ///
  void CEF_CALLBACK  on_focused_node_changed(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      struct _cef_domnode_t* node){
        DEBUG_CALLBACK("on_focused_node_changed\n");
      }

  ///
  // Called when a new message is received from a different process. Return true
  // (1) if the message was handled or false (0) otherwise. It is safe to keep a
  // reference to |message| outside of this callback.
  ///
  int CEF_CALLBACK on_process_message_received_for_render(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      cef_process_id_t source_process,
      struct _cef_process_message_t* message){
        DEBUG_CALLBACK("on_process_message_received_for_render\n");
        return 0;
      }


void initialize_cef_render_process_handler(cef_render_process_handler_t *handler){
    DEBUG_CALLBACK("initialize_cef_render_process_handler\n");
    handler->base.size = sizeof(cef_render_process_handler_t);
    initialize_cef_base_ref_counted((cef_base_ref_counted_t*)handler);
    handler->on_browser_created= on_browser_created;
    handler->on_browser_destroyed= on_browser_destroyed;
    handler->on_context_created= on_context_created;
    handler->on_context_released= on_context_released;
    handler->on_uncaught_exception= on_uncaught_exception;
    handler->on_focused_node_changed= on_focused_node_changed;
    handler->on_process_message_received= on_process_message_received_for_render;
}