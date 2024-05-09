#pragma once

#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_display_handler_capi.h"
#include "utils.h"

///
  /// Called when a frame's address has changed.
  ///
  void CEF_CALLBACK on_address_change(struct _cef_display_handler_t* self,
                                        struct _cef_browser_t* browser,
                                        struct _cef_frame_t* frame,
                                        const cef_string_t* url){
    DEBUG_CALLBACK("on_address_change\n");
  }

  ///
  /// Called when the page title changes.
  ///
  void CEF_CALLBACK on_title_change(struct _cef_display_handler_t* self,
                                      struct _cef_browser_t* browser,
                                      const cef_string_t* title){
    DEBUG_CALLBACK("on_title_change\n");
  }

  ///
  /// Called when the page icon changes.
  ///
  void CEF_CALLBACK on_favicon_urlchange(struct _cef_display_handler_t* self,
                                           struct _cef_browser_t* browser,
                                           cef_string_list_t icon_urls){
      DEBUG_CALLBACK("on_favicon_urlchange\n");
  }

  ///
  /// Called when web content in the page has toggled fullscreen mode. If
  /// |fullscreen| is true (1) the content will automatically be sized to fill
  /// the browser content area. If |fullscreen| is false (0) the content will
  /// automatically return to its original size and position. With the Alloy
  /// runtime the client is responsible for triggering the fullscreen transition
  /// (for example, by calling cef_window_t::SetFullscreen when using Views).
  /// With the Chrome runtime the fullscreen transition will be triggered
  /// automatically. The cef_window_delegate_t::OnWindowFullscreenTransition
  /// function will be called during the fullscreen transition for notification
  /// purposes.
  ///
  void CEF_CALLBACK on_fullscreen_mode_change(
      struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser,
      int fullscreen){
        DEBUG_CALLBACK("on_fullscreen_mode_change\n");
      }

  ///
  /// Called when the browser is about to display a tooltip. |text| contains the
  /// text that will be displayed in the tooltip. To handle the display of the
  /// tooltip yourself return true (1). Otherwise, you can optionally modify
  /// |text| and then return false (0) to allow the browser to display the
  /// tooltip. When window rendering is disabled the application is responsible
  /// for drawing tooltips and the return value is ignored.
  ///
  int CEF_CALLBACK on_tooltip(struct _cef_display_handler_t* self,
                                struct _cef_browser_t* browser,
                                cef_string_t* text){
    DEBUG_CALLBACK("on_tooltip\n");
    return 0;
  }

  ///
  /// Called when the browser receives a status message. |value| contains the
  /// text that will be displayed in the status message.
  ///
  void CEF_CALLBACK on_status_message(struct _cef_display_handler_t* self,
                                        struct _cef_browser_t* browser,
                                        const cef_string_t* value){
    DEBUG_CALLBACK("on_status_message\n");
  }

  ///
  /// Called to display a console message. Return true (1) to stop the message
  /// from being output to the console.
  ///
  int CEF_CALLBACK on_console_message(struct _cef_display_handler_t* self,
                                        struct _cef_browser_t* browser,
                                        cef_log_severity_t level,
                                        const cef_string_t* message,
                                        const cef_string_t* source,
                                        int line){
    DEBUG_CALLBACK("on_console_message\n");
    return 0;
  }

  ///
  /// Called when auto-resize is enabled via
  /// cef_browser_host_t::SetAutoResizeEnabled and the contents have auto-
  /// resized. |new_size| will be the desired size in view coordinates. Return
  /// true (1) if the resize was handled or false (0) for default handling.
  ///
  int CEF_CALLBACK  on_auto_resize(struct _cef_display_handler_t* self,
                                    struct _cef_browser_t* browser,
                                    const cef_size_t* new_size){
    DEBUG_CALLBACK("on_auto_resize\n");
    return 0;
  }

  ///
  /// Called when the overall page loading progress has changed. |progress|
  /// ranges from 0.0 to 1.0.
  ///
  void CEF_CALLBACK on_loading_progress_change(
      struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser,
      double progress){
      DEBUG_CALLBACK("on_loading_progress_change\n");
  }

  ///
  /// Called when the browser's cursor has changed. If |type| is CT_CUSTOM then
  /// |custom_cursor_info| will be populated with the custom cursor information.
  /// Return true (1) if the cursor change was handled or false (0) for default
  /// handling.
  ///
  int CEF_CALLBACK on_cursor_change(
      struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser,
      cef_cursor_handle_t cursor,
      cef_cursor_type_t type,
      const cef_cursor_info_t* custom_cursor_info){
        DEBUG_CALLBACK("on_cursor_change\n");
        return 0;
  }

  ///
  /// Called when the browser's access to an audio and/or video source has
  /// changed.
  ///
  void CEF_CALLBACK on_media_access_change(
      struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser,
      int has_video_access,
      int has_audio_access){
        DEBUG_CALLBACK("on_media_access_change\n");
  }



display_handler *initialize_display_handler(){
    DEBUG_CALLBACK("initialize_display_handler\n");
    display_handler *l=calloc(1,sizeof(display_handler));
    cef_display_handler_t *handler = (cef_display_handler_t *)l;
    handler->on_address_change= on_address_change;
    handler->on_auto_resize= on_auto_resize;
    handler->on_console_message=on_console_message;
    handler->on_cursor_change=on_cursor_change;
    handler->on_favicon_urlchange=on_favicon_urlchange;
    handler->on_fullscreen_mode_change=on_fullscreen_mode_change;
    handler->on_loading_progress_change=on_loading_progress_change;
    handler->on_media_access_change=on_media_access_change;
    handler->on_status_message=on_status_message;
    handler->on_title_change=on_title_change;
    handler->on_tooltip=on_tooltip;
    return l;
}


void initialize_display_handler_direct(cef_display_handler_t *handler){
    DEBUG_CALLBACK("initialize_display_handler\n");
    handler->base.size = sizeof(cef_display_handler_t);
    initialize_cef_base_ref_counted((cef_base_ref_counted_t*)handler);
    handler->on_address_change= on_address_change;
    handler->on_auto_resize= on_auto_resize;
    handler->on_console_message=on_console_message;
    handler->on_cursor_change=on_cursor_change;
    handler->on_favicon_urlchange=on_favicon_urlchange;
    handler->on_fullscreen_mode_change=on_fullscreen_mode_change;
    handler->on_loading_progress_change=on_loading_progress_change;
    handler->on_media_access_change=on_media_access_change;
    handler->on_status_message=on_status_message;
    handler->on_title_change=on_title_change;
    handler->on_tooltip=on_tooltip;
}