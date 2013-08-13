package gofish
/*
#include <string.h>
#include <stdlib.h>

extern int decrypt_string(const char *key, const char *str, char *dest, int len);
extern int encrypt_string(const char *key, const char *str, char *dest, int len);

*/
import "C"
import (
    "unsafe"
)

func Decrypt(key string, message string) string {
    ckey := C.CString(key)
    defer C.free(unsafe.Pointer(ckey))
    cstr := C.CString(message)
    defer C.free(unsafe.Pointer(cstr))
    cdest := C.CString("")
    defer C.free(unsafe.Pointer(cdest))
    clen := C.int(len(message))
    C.decrypt_string(ckey, cstr, cdest, clen)
    ret := C.GoString(cdest)
    return ret
} 

func Encrypt(key string, message string) string {
    ckey := C.CString(key)
    defer C.free(unsafe.Pointer(ckey))
    cstr := C.CString(message)
    defer C.free(unsafe.Pointer(cstr))
    cdest := C.CString("")
    defer C.free(unsafe.Pointer(cdest))
    clen := C.int(len(message))
    C.encrypt_string(ckey, cstr, cdest, clen)
    ret := C.GoString(cdest)
    return ret
}
