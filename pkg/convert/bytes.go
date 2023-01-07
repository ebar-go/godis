package convert

import (
	"reflect"
	"unsafe"
)

func String2Byte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	slice := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&slice))
}

func Byte2String(p []byte) string {
	return *(*string)(unsafe.Pointer(&p))
}

func ToByte(val any) []byte {
	switch val.(type) {
	case []byte:
		return val.([]byte)
	case string:
		return String2Byte(val.(string))
	}

	return nil
}
