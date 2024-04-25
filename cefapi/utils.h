
#include <string.h>

cef_string_t getCefString(const char *utf8String){
    cef_string_t cefString = {};
    cef_string_utf8_to_utf16(utf8String, strlen(utf8String),&cefString);
    return cefString;
}