// CEF C API example
// Project website: https://github.com/cztomczak/cefcapi

#pragma once

#include <stdatomic.h>
#include "include/capi/cef_base_capi.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/views/cef_window_capi.h"
#include "include/capi/cef_display_handler_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
#include "include/capi/cef_drag_handler_capi.h"
#include "include/capi/cef_load_handler_capi.h"
#include "include/capi/cef_render_process_handler_capi.h"


// Set to 1 to check if add_ref() and release()
// are called and to track the total number of calls.
// add_ref will be printed as "+", release as "-".
#define DEBUG_REFERENCE_COUNTING 0

// Print only the first execution of the callback,
// ignore the subsequent.
// #define DEBUG_CALLBACK(x) { \
//     static int first_call = 1; \
//     if (first_call == 1) { \
//         first_call = 0; \
//         printf(x); \
//     } \
// }

#define DEBUG_CALLBACK(x) printf(x);
// ----------------------------------------------------------------------------
// cef_base_ref_counted_t
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// cef_base_ref_counted_t
// ----------------------------------------------------------------------------

///
// Structure defining the reference count implementation functions.
// All framework structures must include the cef_base_ref_counted_t
// structure first.
///

///
// Increment the reference count.
///
#define ADD_REF(type) \
static \
void \
CEF_CALLBACK \
type##_add_ref(cef_base_ref_counted_t *self) \
{ \
	struct _##type *handler = (struct _##type *)self; \
	/*printf("process %d add ref " #type " %d->%d \n",GetCurrentProcessId(),handler->ref_count,handler->ref_count+1); \*/ \
	atomic_fetch_add(&handler->ref_count, 1); \
}

///
// Decrement the reference count.  Delete this object when no references
// remain.
///
#define RELEASE(type) \
static \
int \
CEF_CALLBACK \
type##_release(cef_base_ref_counted_t* self) { \
	struct _##type *handler = (struct _##type *)self; \
	/*printf("process %d release ref " #type " %d->%d \n",GetCurrentProcessId(),handler->ref_count,handler->ref_count-1); \ */ \
	if (atomic_fetch_sub(&handler->ref_count, 1) - 1 == 0) { \
		free(handler); \
		return 1; \
	} \
	return 0; \
}

///
// Returns the current number of references.
///
#define HAS_ONE_REF(type) \
static \
int \
CEF_CALLBACK \
type##_has_one_ref(cef_base_ref_counted_t* self) { \
	struct _##type *handler = (struct _##type *)self; \
	return atomic_load(&handler->ref_count) == 1; \
}

#define IMPLEMENT_REFCOUNTING(type) \
ADD_REF(type) \
RELEASE(type) \
HAS_ONE_REF(type)

#define GENERATE_CEF_BASE_INITIALIZER(type) \
void \
initialize_##type##_base(type *object) \
{ \
	cef_base_ref_counted_t *base = (cef_base_ref_counted_t *)object; \
	base->size = sizeof(type); \
	base->add_ref = type##_add_ref; \
	base->release = type##_release; \
	base->has_one_ref = type##_has_one_ref; \
}

struct _life_span_handler_t;
struct _client_t;
struct _app_t;
struct _invocation_handler;
struct _drag_handler;
struct _display_handler;
struct _view_delegate_t;
struct _browser_view_delegate_t;
struct _load_handler;
struct _render_handler;
struct _render_process_handler;
struct _resource_handler;

void initialize_life_span_handler_t_base(struct _life_span_handler_t *object);
void initialize_view_delegate_t_base(struct _view_delegate_t *object);
void initialize_browser_view_delegate_t_base(struct _browser_view_delegate_t *object);
void initialize_drag_handler_base(struct _drag_handler *object);
void initialize_display_handler_base(struct _display_handler *object);
void initialize_client_t_base(struct _client_t *object);
void initialize_app_t_base(struct _app_t *object);
void initialize_invocation_handler_base(struct _invocation_handler *object);
void initialize_load_handler_base(struct _load_handler *object);
void initialize_render_handler_base(struct _render_handler *object);
void initialize_render_process_handler_base(struct _render_process_handler *object);
void initialize_resource_handler_base(struct _resource_handler *object);

#define initialize_cef_base(T) \
    _Generic((T), \
	struct _life_span_handler_t*: initialize_life_span_handler_t_base, \
	struct _view_delegate_t*: initialize_view_delegate_t_base, \
	struct _browser_view_delegate_t*: initialize_browser_view_delegate_t_base, \
	struct _drag_handler*: initialize_drag_handler_base, \
	struct _display_handler*: initialize_display_handler_base, \
	struct _client_t*: initialize_client_t_base, \
	struct _app_t*: initialize_app_t_base, \
	struct _load_handler*: initialize_load_handler_base, \
	struct _render_handler*: initialize_render_handler_base, \
	struct _render_process_handler*: initialize_render_process_handler_base, \
	struct _resource_handler*: initialize_resource_handler_base, \
	struct _invocation_handler*: initialize_invocation_handler_base)(T)



typedef struct _app_t {
	cef_app_t app;
	atomic_int ref_count;
} app_t;

IMPLEMENT_REFCOUNTING(app_t)
GENERATE_CEF_BASE_INITIALIZER(app_t)


typedef struct _client_t {
	cef_client_t client;
	cef_window_t* window;
	atomic_int ref_count;
} client_t;

IMPLEMENT_REFCOUNTING(client_t)
GENERATE_CEF_BASE_INITIALIZER(client_t)


typedef struct _display_handler {
    cef_display_handler_t handler;
    atomic_int ref_count;
} display_handler;

IMPLEMENT_REFCOUNTING(display_handler)
GENERATE_CEF_BASE_INITIALIZER(display_handler)


typedef struct _drag_handler {
	cef_drag_handler_t handler;
	atomic_int ref_count;
} drag_handler;

IMPLEMENT_REFCOUNTING(drag_handler)
GENERATE_CEF_BASE_INITIALIZER(drag_handler)


typedef struct _life_span_handler_t {
	cef_life_span_handler_t handler;
	atomic_int ref_count;
} life_span_handler_t;

IMPLEMENT_REFCOUNTING(life_span_handler_t)
GENERATE_CEF_BASE_INITIALIZER(life_span_handler_t)


typedef struct _invocation_handler {
	cef_v8handler_t handler;
	atomic_int ref_count;
} invocation_handler;

IMPLEMENT_REFCOUNTING(invocation_handler)
GENERATE_CEF_BASE_INITIALIZER(invocation_handler)

typedef struct _load_handler {
	cef_load_handler_t handler;
	atomic_int ref_count;
} load_handler;

IMPLEMENT_REFCOUNTING(load_handler)
GENERATE_CEF_BASE_INITIALIZER(load_handler)

typedef struct _render_handler {
	cef_render_handler_t handler;
	atomic_int ref_count;
} render_handler;

IMPLEMENT_REFCOUNTING(render_handler)
GENERATE_CEF_BASE_INITIALIZER(render_handler)

typedef struct _render_process_handler {
	cef_render_process_handler_t handler;
	atomic_int ref_count;
} render_process_handler;

IMPLEMENT_REFCOUNTING(render_process_handler)
GENERATE_CEF_BASE_INITIALIZER(render_process_handler)


typedef struct _resource_handler {
	cef_resource_handler_t handler;
	atomic_int ref_count;
	int request_id;
} resource_handler;
IMPLEMENT_REFCOUNTING(resource_handler)
GENERATE_CEF_BASE_INITIALIZER(resource_handler)

///
// Increment the reference count.
///
void CEF_CALLBACK add_ref(cef_base_ref_counted_t* self) {
    //DEBUG_CALLBACK("cef_base_ref_counted_t.add_ref\n");
    if (DEBUG_REFERENCE_COUNTING)
        printf("+");
}

///
// Decrement the reference count.  Delete this object when no references
// remain.
///
int CEF_CALLBACK release(cef_base_ref_counted_t* self) {
    //DEBUG_CALLBACK("cef_base_ref_counted_t.release\n");
    if (DEBUG_REFERENCE_COUNTING)
        printf("-");
    return 1;
}

///
// Returns the current number of references.
///
int CEF_CALLBACK has_one_ref(cef_base_ref_counted_t* self) {
    DEBUG_CALLBACK("cef_base_ref_counted_t.has_one_ref\n");
    if (DEBUG_REFERENCE_COUNTING)
        printf("=");
    return 1;
}

void initialize_cef_base_ref_counted(cef_base_ref_counted_t* base) {
    printf("initialize_cef_base_ref_counted\n");
    // Check if "size" member was set.
    size_t size = base->size;
    // Let's print the size in case sizeof was used
    // on a pointer instead of a structure. In such
    // case the number will be very high.
    printf("cef_base_ref_counted_t.size = %lu\n", (unsigned long)size);
    if (size <= 0) {
        printf("FATAL: initialize_cef_base failed, size member not set\n");
        _exit(1);
    }
    base->add_ref = add_ref;
    base->release = release;
    base->has_one_ref = has_one_ref;
}
