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
        if (isEqual(name,"testfunc")) {
            DEBUG_CALLBACK("v8_execute testfunc\n");
            cef_string_t resultData=getCefString("test func value");
            cef_v8value_t *result= cef_v8value_create_string(&resultData);
            *retval = result;
            return 1;
        }
        return 0;
    }


invocation_handler* initialize_v8handler(){
    DEBUG_CALLBACK("initialize_v8handler\n");
    invocation_handler  *h;
    h = calloc(1, sizeof(invocation_handler));
    initialize_cef_base(h);
    cef_v8handler_t *handler = (cef_v8handler_t *)h;
    handler->base.add_ref((cef_base_ref_counted_t *)h);
    handler->execute = v8_execute;
    return h;
}