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

    buf := make([]byte, 1000)

    clen := C.int(len(message))

    C.decrypt_string(ckey, cstr, (*C.char)(unsafe.Pointer(&buf[0])), clen)

    return string(buf)
} 

func Encrypt(key string, message string) string {
    ckey := C.CString(key)
    defer C.free(unsafe.Pointer(ckey))

    cstr := C.CString(message)
    defer C.free(unsafe.Pointer(cstr))

    buf := make([]byte, 1000)

    clen := C.int(len(message))

    C.encrypt_string(ckey, cstr, (*C.char)(unsafe.Pointer(&buf[0])), clen)

    return "+OK " + string(buf)
}
