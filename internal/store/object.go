package store

import (
	"github.com/ebar-go/godis/pkg/convert"
	"github.com/ebar-go/godis/server/types"
	"strconv"
	"unsafe"
)

type ObjectType uint

const (
	ObjectString = iota
	ObjectList
	ObjectHash
	ObjectSet
	ObjectSortedSet
)

type ObjectEncoding uint

const (
	EncodingRaw = iota
	EncodingInt
	EncodingHT
	EncodingZipMap
	EncodingLinkedList
	EncodingZipList
	EncodingIntSet
	EncodingSkipList
	EncodingEMBStr
	EncodingQuickList
	EncodingListPack
)

type Object struct {
	Type     ObjectType     // 标识该对象是什么类型的对象
	Encoding ObjectEncoding // 底层的数据结构
	Ptr      unsafe.Pointer // 指向底层数据结构的指针
}

func (obj Object) Len() uint64 {
	switch obj.Type {
	case ObjectString:
		sds := (*types.SDS)(obj.Ptr)
		return sds.Len()
	}

	return 0
}

func (obj *Object) SetStringValue(val any) {
	// 如果是字符串
	if obj.Encoding == EncodingRaw {
		switch val.(type) {
		case string:
			sds := (*types.SDS)(obj.Ptr)
			sds.Set(convert.ToByte(val))
		case int:
			obj.Encoding = EncodingInt
			obj.Ptr = unsafe.Pointer(&val)
		}
	} else if obj.Encoding == EncodingInt { // 如果是数字
		switch val.(type) {
		case string:
			obj.Encoding = EncodingRaw // 切换编码
			obj.Ptr = unsafe.Pointer(types.NewSDS(convert.String2Byte(val.(string))))
		case int:
			obj.Ptr = unsafe.Pointer(&val)
		}
	}

}

func (obj Object) String() string {
	switch obj.Type {
	case ObjectString:
		if obj.Encoding == EncodingRaw {
			sds := (*types.SDS)(obj.Ptr)
			return sds.String()
		} else if obj.Encoding == EncodingInt {
			return strconv.Itoa(*(*int)(obj.Ptr))
		}

	}

	return ""
}

func NewKeyObject(key string) *Object {
	return NewStringObjectWithEncoding(key, EncodingRaw)
}

func NewStringObject(val any) *Object {
	switch val.(type) {
	case string:
		return NewStringObjectWithEncoding(val, EncodingRaw)
	case int:
		return NewStringObjectWithEncoding(val, EncodingInt)
	}

	return nil
}

func NewStringObjectWithEncoding(val any, encoding ObjectEncoding) *Object {
	obj := &Object{Type: ObjectString, Encoding: encoding}
	switch encoding {
	case EncodingRaw:
		obj.Ptr = unsafe.Pointer(types.NewSDS(convert.ToByte(val)))
	case EncodingInt:
		n, _ := val.(int)
		obj.Ptr = unsafe.Pointer(&n)
	}
	return obj
}