package convert

import (
	"encoding/binary"
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
	case int64:
		var buf = make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(val.(int64)))
		return buf
	}

	return nil
}
