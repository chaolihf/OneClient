#pragma once

#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_resource_handler_capi.h"
#include "cef_v8handler.h"
#include "utils.h"

  char* content_data_ =
        "<html><head><title>Scheme Test(Home Page)</title></head>"
        "<body bgcolor=\"white\">"
        "This contents of this page page are served by the resource handlle , navigate to new page <a href='https://www.sina.com.cn'>Sina</></body></html>";
  size_t offset_ = 0;

  ///
  /// Open the response stream. To handle the request immediately set
  /// |handle_request| to true (1) and return true (1). To decide at a later
  /// time set |handle_request| to false (0), return true (1), and execute
  /// |callback| to continue or cancel the request. To cancel the request
  /// immediately set |handle_request| to true (1) and return false (0). This
  /// function will be called in sequence but not from a dedicated thread. For
  /// backwards compatibility set |handle_request| to false (0) and return false
  /// (0) and the ProcessRequest function will be called.
  ///
  int CEF_CALLBACK resource_handler_open(struct _cef_resource_handler_t* self,
                          struct _cef_request_t* request,
                          int* handle_request,
                          struct _cef_callback_t* callback){
    DEBUG_CALLBACK("resource_handler_open\n");
    *handle_request=1;
    const cef_string_userfree_t url=request->get_url(request);
    cef_string_userfree_free(url);
    return 1;
  }

  ///
  /// Retrieve response header information. If the response length is not known
  /// set |response_length| to -1 and read_response() will be called until it
  /// returns false (0). If the response length is known set |response_length|
  /// to a positive value and read_response() will be called until it returns
  /// false (0) or the specified number of bytes have been read. Use the
  /// |response| object to set the mime type, http status code and other
  /// optional header values. To redirect the request to a new URL set
  /// |redirectUrl| to the new URL. |redirectUrl| can be either a relative or
  /// fully qualified URL. It is also possible to set |response| to a redirect
  /// http status code and pass the new URL via a Location header. Likewise with
  /// |redirectUrl| it is valid to set a relative or fully qualified URL as the
  /// Location header value. If an error occured while setting up the request
  /// you can call set_error() on |response| to indicate the error condition.
  ///
  void CEF_CALLBACK get_response_headers(struct _cef_resource_handler_t* self,
                                           struct _cef_response_t* response,
                                           int64_t* response_length,
                                           cef_string_t* redirectUrl){
    DEBUG_CALLBACK("resource handler get_response_headers\n");
    cef_string_t mineType=getCefString("text/html");
    response->set_mime_type(response,&mineType);
    response->set_status(response,200);
    *response_length = strlen(content_data_);
  }

  ///
  /// Skip response data when requested by a Range header. Skip over and discard
  /// |bytes_to_skip| bytes of response data. If data is available immediately
  /// set |bytes_skipped| to the number of bytes skipped and return true (1). To
  /// read the data at a later time set |bytes_skipped| to 0, return true (1)
  /// and execute |callback| when the data is available. To indicate failure set
  /// |bytes_skipped| to < 0 (e.g. -2 for ERR_FAILED) and return false (0). This
  /// function will be called in sequence but not from a dedicated thread.
  ///
  int CEF_CALLBACK resource_handler_skip(struct _cef_resource_handler_t* self,
                          int64_t bytes_to_skip,
                          int64_t* bytes_skipped,
                          struct _cef_resource_skip_callback_t* callback){
    DEBUG_CALLBACK("resource_handler_skip\n");
    *bytes_skipped=0;
    return 1;
  }

  ///
  /// Read response data. If data is available immediately copy up to
  /// |bytes_to_read| bytes into |data_out|, set |bytes_read| to the number of
  /// bytes copied, and return true (1). To read the data at a later time keep a
  /// pointer to |data_out|, set |bytes_read| to 0, return true (1) and execute
  /// |callback| when the data is available (|data_out| will remain valid until
  /// the callback is executed). To indicate response completion set
  /// |bytes_read| to 0 and return false (0). To indicate failure set
  /// |bytes_read| to < 0 (e.g. -2 for ERR_FAILED) and return false (0). This
  /// function will be called in sequence but not from a dedicated thread. For
  /// backwards compatibility set |bytes_read| to -1 and return false (0) and
  /// the ReadResponse function will be called.
  ///
  int CEF_CALLBACK resource_handler_read(struct _cef_resource_handler_t* self,
                          void* data_out,
                          int bytes_to_read,
                          int* bytes_read,
                          struct _cef_resource_read_callback_t* callback){
    DEBUG_CALLBACK("resource_handler_read\n");
    bool has_data = false;
    *bytes_read = 0;

    if (offset_ < strlen(content_data_)) {
      // Copy the next block of data into the buffer.
      int transfer_size =min(bytes_to_read, strlen(content_data_) - offset_);
      memcpy(data_out, content_data_ + offset_, transfer_size);
      offset_ += transfer_size;

      *bytes_read = transfer_size;
      has_data = true;
    }
    return has_data;
  }

  ///
  /// Request processing has been canceled.
  ///
  void CEF_CALLBACK  resource_handler_cancel(struct _cef_resource_handler_t* self){
    DEBUG_CALLBACK("resource_handler_cancel\n");
  }

void initialize_cef_resource_handler_direct(cef_resource_handler_t *handler){
    DEBUG_CALLBACK("initialize_cef_resource_handler_direct\n");
    handler->base.size = sizeof(cef_resource_handler_t);
    initialize_cef_base_ref_counted((cef_base_ref_counted_t*)handler);
    handler->cancel=resource_handler_cancel;
    handler->read=resource_handler_read;
    handler->skip=resource_handler_skip;
    handler->get_response_headers=get_response_headers;
    handler->open=resource_handler_open;
}