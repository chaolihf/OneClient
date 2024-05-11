#pragma once

#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_scheme_capi.h"
#include "cef_v8handler.h"
#include "utils.h"
extern cef_resource_handler_t g_resource_handler;

///
  /// Return a new resource handler instance to handle the request or an NULL
  /// reference to allow default handling of the request. |browser| and |frame|
  /// will be the browser window and frame respectively that originated the
  /// request or NULL if the request did not originate from a browser window
  /// (for example, if the request came from cef_urlrequest_t). The |request|
  /// object passed to this function cannot be modified.
  ///
  struct _cef_resource_handler_t* CEF_CALLBACK create(
      struct _cef_scheme_handler_factory_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      const cef_string_t* scheme_name,
      struct _cef_request_t* request){
        DEBUG_CALLBACK("create resource handler\n");
        cef_string_userfree_t url=request->get_url(request);
        cef_string_t localUrl=getCefString("http://myhomepage/");
        if(cef_string_utf16_cmp(url,&localUrl)==0){
          cef_string_userfree_free(url);
          return &g_resource_handler; 
        }
        else{
          cef_string_userfree_free(url);
          return NULL;
        }
  }

    
void initialize_cef_schema_handler_factory_direct(cef_scheme_handler_factory_t *handler){
    DEBUG_CALLBACK("initialize_cef_schema_handler_factory_direct\n");
    handler->base.size = sizeof(cef_scheme_handler_factory_t);
    initialize_cef_base_ref_counted((cef_base_ref_counted_t*)handler);
    handler->create=create;
}