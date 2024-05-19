#ifndef COMMON_UTILS_H
#define COMMON_UTILS_H
#include <string.h>
#include "include/capi/cef_base_capi.h"
#include <stdio.h>
#include "cefapi.h"


cef_string_t getCefString(const char *utf8String){
    cef_string_t cefString = {};
    cef_string_utf8_to_utf16(utf8String, strlen(utf8String),&cefString);
    return cefString;
}

cef_string_t getCefStringFromGo(const Go_String *source){
    cef_string_t cefString = {};
    cef_string_utf8_to_utf16(source->p, source->n,&cefString);
    return cefString;
}

bool isEqual(const cef_string_t *source, char *compare){
    cef_string_t newCompare=getCefString(compare);
    return cef_string_utf16_cmp(source,&newCompare)==0;
}

char * convertCefStringToChar(const cef_string_t *source) {
    cef_string_utf8_t cefString={};
    cef_string_utf16_to_utf8(source->str,source->length,&cefString);
    char* result=(char *)calloc(1,cefString.length+1);
    memcpy(result,cefString.str,cefString.length);
    result[cefString.length]='\0';
    return result;
}
// char* ConvertCefStringToChar(cef_string_userfree_t source) {
    
//   if (source && source->length > 0 && source->str) {
//     cef_string_userfree_wide_t wide_str = {};
//     cef_string_utf16_to_wide(source->str, source->length, wide_str);
//     wide_str.
//     cef_string_userfree_free(source);
    
//     // 获取source的长度（以字符为单位）
//     size_t length = (size_t)source->length;

//     // 分配足够大小的char*缓冲区
//     char* result = (char*)malloc((length + 1) * sizeof(char));

//     // 将cef_string复制到char*缓冲区中
//     strncpy(result, source->str, length);
//     result[length] = '\0'; // 确保以null终止

//     return result;
//   }

//   return NULL;
// }
#endif