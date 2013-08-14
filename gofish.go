package gofish
/*
#include <string.h>
#include <stdlib.h>

extern int decrypt_string(const char *key, const char *str, char *dest, int len);
extern int encrypt_string(const char *key, const char *str, char *dest, int len);

*/
import "C"
import "unsafe"

func Decrypt(key string, message string) string {
    ckey := C.CString(key)
    defer C.free(unsafe.Pointer(ckey))

    cstr := C.CString(message)
    defer C.free(unsafe.Pointer(cstr))

    buf := make([]byte, 1000)

    clen := C.int(len(message))

    len := C.decrypt_string(ckey, cstr, (*C.char)(unsafe.Pointer(&buf[0])), clen)

    buf = buf[0:int(len)]

    return string(buf)
} 

func Encrypt(key string, message string) string {
    ckey := C.CString(key)
    defer C.free(unsafe.Pointer(ckey))

    cstr := C.CString(message)
    defer C.free(unsafe.Pointer(cstr))

    buf := make([]byte, 1000)

    clen := C.int(len(message))

    len := C.encrypt_string(ckey, cstr, (*C.char)(unsafe.Pointer(&buf[0])), clen)
    buf = buf[0:int(len)]

    return "+OK " + string(buf)
}
