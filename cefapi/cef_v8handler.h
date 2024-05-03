#pragma once
#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "utils.h"



  ///
  // Handle execution of the function identified by |name|. |object| is the
  // receiver ('this' object) of the function. |arguments| is the list of
  // arguments passed to the function. If execution succeeds set |retval| to the
  // function return value. If execution fails set |exception| to the exception
  // that will be thrown. Return true (1) if execution was handled.
  ///
  int CEF_CALLBACK v8_execute(struct _cef_v8handler_t* self,
                             const cef_string_t* name,
                             struct _cef_v8value_t* object,
                             size_t argumentsCount,
                             struct _cef_v8value_t* const* arguments,
                             struct _cef_v8value_t** retval,
                             cef_string_t* exception){
        DEBUG_CALLBACK("v8_execute\n");
        return 1;
    }


void initialize_v8handler(cef_v8handler_t *handler){
    DEBUG_CALLBACK("initialize_v8handler\n");
    handler->base.size = sizeof(cef_v8handler_t);
    initialize_cef_base_ref_counted((cef_base_ref_counted_t*)handler);
    handler->execute = v8_execute;
}