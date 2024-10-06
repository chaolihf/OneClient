#pragma once

#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_render_handler_capi.h"
#include "utils.h"

///
  /// Return the handler for accessibility notifications. If no handler is
  /// provided the default implementation will be used.
  ///
  struct _cef_accessibility_handler_t* CEF_CALLBACK get_accessibility_handler(
      struct _cef_render_handler_t* self){
    DEBUG_CALLBACK("get_accessibility_handler\n");
  }

  ///
  /// Called to retrieve the root window rectangle in screen DIP coordinates.
  /// Return true (1) if the rectangle was provided. If this function returns
  /// false (0) the rectangle from GetViewRect will be used.
  ///
  int CEF_CALLBACK get_root_screen_rect(struct _cef_render_handler_t* self,
                                          struct _cef_browser_t* browser,
                                          cef_rect_t* rect){
     DEBUG_CALLBACK("get_root_screen_rect\n");
     return 0;
  }

  ///
  /// Called to retrieve the view rectangle in screen DIP coordinates. This
  /// function must always provide a non-NULL rectangle.
  ///
  void CEF_CALLBACK get_view_rect(struct _cef_render_handler_t* self,
                                    struct _cef_browser_t* browser,
                                    cef_rect_t* rect){
    DEBUG_CALLBACK("get_view_rect\n");

  }

  ///
  /// Called to retrieve the translation from view DIP coordinates to screen
  /// coordinates. Windows/Linux should provide screen device (pixel)
  /// coordinates and MacOS should provide screen DIP coordinates. Return true
  /// (1) if the requested coordinates were provided.
  ///
  int CEF_CALLBACK get_screen_point(struct _cef_render_handler_t* self,
                                      struct _cef_browser_t* browser,
                                      int viewX,
                                      int viewY,
                                      int* screenX,
                                      int* screenY){
    DEBUG_CALLBACK("get_screen_point\n");

 }

  ///
  /// Called to allow the client to fill in the CefScreenInfo object with
  /// appropriate values. Return true (1) if the |screen_info| structure has
  /// been modified.
  ///
  /// If the screen info rectangle is left NULL the rectangle from GetViewRect
  /// will be used. If the rectangle is still NULL or invalid popups may not be
  /// drawn correctly.
  ///
  int CEF_CALLBACK get_screen_info(struct _cef_render_handler_t* self,
                                     struct _cef_browser_t* browser,
                                     cef_screen_info_t* screen_info){
     DEBUG_CALLBACK("get_screen_info\n");
     return 0;
  }

  ///
  /// Called when the browser wants to show or hide the popup widget. The popup
  /// should be shown if |show| is true (1) and hidden if |show| is false (0).
  ///
  void CEF_CALLBACK on_popup_show(struct _cef_render_handler_t* self,
                                    struct _cef_browser_t* browser,
                                    int show){
    DEBUG_CALLBACK("on_popup_show\n");

  }

  ///
  /// Called when the browser wants to move or resize the popup widget. |rect|
  /// contains the new location and size in view coordinates.
  ///
  void CEF_CALLBACK on_popup_size(struct _cef_render_handler_t* self,
                                    struct _cef_browser_t* browser,
                                    const cef_rect_t* rect){
    DEBUG_CALLBACK("on_popup_size\n");

  }

  ///
  /// Called when an element should be painted. Pixel values passed to this
  /// function are scaled relative to view coordinates based on the value of
  /// CefScreenInfo.device_scale_factor returned from GetScreenInfo. |type|
  /// indicates whether the element is the view or the popup widget. |buffer|
  /// contains the pixel data for the whole image. |dirtyRects| contains the set
  /// of rectangles in pixel coordinates that need to be repainted. |buffer|
  /// will be |width|*|height|*4 bytes in size and represents a BGRA image with
  /// an upper-left origin. This function is only called when
  /// cef_window_tInfo::shared_texture_enabled is set to false (0).
  ///
  void CEF_CALLBACK on_paint(struct _cef_render_handler_t* self,
                               struct _cef_browser_t* browser,
                               cef_paint_element_type_t type,
                               size_t dirtyRectsCount,
                               cef_rect_t const* dirtyRects,
                               const void* buffer,
                               int width,
                               int height){
    DEBUG_CALLBACK("on_paint\n");

  }

  ///
  /// Called when an element has been rendered to the shared texture handle.
  /// |type| indicates whether the element is the view or the popup widget.
  /// |dirtyRects| contains the set of rectangles in pixel coordinates that need
  /// to be repainted. |shared_handle| is the handle for a D3D11 Texture2D that
  /// can be accessed via ID3D11Device using the OpenSharedResource function.
  /// This function is only called when cef_window_tInfo::shared_texture_enabled
  /// is set to true (1), and is currently only supported on Windows.
  ///
  void CEF_CALLBACK on_accelerated_paint(struct _cef_render_handler_t* self,
                                           struct _cef_browser_t* browser,
                                           cef_paint_element_type_t type,
                                           size_t dirtyRectsCount,
                                           cef_rect_t const* dirtyRects,
                                           void* shared_handle){
    DEBUG_CALLBACK("on_accelerated_paint\n");

  }

  ///
  /// Called to retrieve the size of the touch handle for the specified
  /// |orientation|.
  ///
  void CEF_CALLBACK get_touch_handle_size(
      struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser,
      cef_horizontal_alignment_t orientation,
      cef_size_t* size){
    DEBUG_CALLBACK("get_touch_handle_size\n");

  }

  ///
  /// Called when touch handle state is updated. The client is responsible for
  /// rendering the touch handles.
  ///
  void CEF_CALLBACK on_touch_handle_state_changed(
      struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser,
      const cef_touch_handle_state_t* state){
    DEBUG_CALLBACK("on_touch_handle_state_changed\n");

  }

  ///
  /// Called when the user starts dragging content in the web view. Contextual
  /// information about the dragged content is supplied by |drag_data|. (|x|,
  /// |y|) is the drag start location in screen coordinates. OS APIs that run a
  /// system message loop may be used within the StartDragging call.
  ///
  /// Return false (0) to abort the drag operation. Don't call any of
  /// cef_browser_host_t::DragSource*Ended* functions after returning false (0).
  ///
  /// Return true (1) to handle the drag operation. Call
  /// cef_browser_host_t::DragSourceEndedAt and DragSourceSystemDragEnded either
  /// synchronously or asynchronously to inform the web view that the drag
  /// operation has ended.
  ///
  int CEF_CALLBACK start_dragging(struct _cef_render_handler_t* self,
                                    struct _cef_browser_t* browser,
                                    struct _cef_drag_data_t* drag_data,
                                    cef_drag_operations_mask_t allowed_ops,
                                    int x,
                                    int y){
     DEBUG_CALLBACK("start_dragging\n");
     return 1;
  }

  ///
  /// Called when the web view wants to update the mouse cursor during a drag &
  /// drop operation. |operation| describes the allowed operation (none, move,
  /// copy, link).
  ///
  void CEF_CALLBACK update_drag_cursor(struct _cef_render_handler_t* self,
                                         struct _cef_browser_t* browser,
                                         cef_drag_operations_mask_t operation){
    DEBUG_CALLBACK("update_drag_cursor\n");

  }

  ///
  /// Called when the scroll offset has changed.
  ///
  void CEF_CALLBACK on_scroll_offset_changed(
      struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser,
      double x,
      double y){
    DEBUG_CALLBACK("on_scroll_offset_changed\n");

  }

  ///
  /// Called when the IME composition range has changed. |selected_range| is the
  /// range of characters that have been selected. |character_bounds| is the
  /// bounds of each character in view coordinates.
  ///
  void CEF_CALLBACK on_ime_composition_range_changed(
      struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser,
      const cef_range_t* selected_range,
      size_t character_boundsCount,
      cef_rect_t const* character_bounds){
    DEBUG_CALLBACK("on_ime_composition_range_changed\n");

  }

  ///
  /// Called when text selection has changed for the specified |browser|.
  /// |selected_text| is the currently selected text and |selected_range| is the
  /// character range.
  ///
  void CEF_CALLBACK on_text_selection_changed(
      struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser,
      const cef_string_t* selected_text,
      const cef_range_t* selected_range){
    DEBUG_CALLBACK("on_text_selection_changed\n");

  }

  ///
  /// Called when an on-screen keyboard should be shown or hidden for the
  /// specified |browser|. |input_mode| specifies what kind of keyboard should
  /// be opened. If |input_mode| is CEF_TEXT_INPUT_MODE_NONE, any existing
  /// keyboard for this browser should be hidden.
  ///
  void CEF_CALLBACK on_virtual_keyboard_requested(
      struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser,
      cef_text_input_mode_t input_mode){
    DEBUG_CALLBACK("on_virtual_keyboard_requested\n");

  }

render_handler *initialize_cef_render_handler(){
    DEBUG_CALLBACK("initialize_cef_render_handler\n");
    render_handler *h=calloc(1,sizeof(render_handler));
    cef_render_handler_t *handler = &h->handler;
    initialize_cef_base(h);
    handler->base.add_ref((cef_base_ref_counted_t *)h);

    handler->get_accessibility_handler= get_accessibility_handler;
    handler->get_root_screen_rect=get_root_screen_rect;
    handler->get_screen_info=get_screen_info;
    handler->get_screen_point=get_screen_point;
    handler->get_touch_handle_size=get_touch_handle_size;
    handler->get_view_rect=get_view_rect;
    handler->on_accelerated_paint=on_accelerated_paint;
    handler->on_ime_composition_range_changed=on_ime_composition_range_changed;
    handler->on_paint=on_paint;
    handler->on_popup_show=on_popup_show;
    handler->on_popup_size=on_popup_size;
    handler->on_scroll_offset_changed=on_scroll_offset_changed;
    handler->on_text_selection_changed=on_text_selection_changed;
    handler->on_touch_handle_state_changed=on_touch_handle_state_changed;
    handler->on_virtual_keyboard_requested=on_virtual_keyboard_requested;
    handler->start_dragging=start_dragging;
    handler->update_drag_cursor=update_drag_cursor;
    return h;
}