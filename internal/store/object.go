package store

import (
	"github.com/ebar-go/godis/errors"
	"github.com/ebar-go/godis/internal/types"
	"github.com/ebar-go/godis/pkg/convert"
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
		case int:
			obj.Encoding = EncodingInt
			obj.Ptr = unsafe.Pointer(&val)
		default:
			sds := (*types.SDS)(obj.Ptr)
			sds.Set(convert.ToByte(val))
		}
	} else if obj.Encoding == EncodingInt { // 如果是数字
		switch val.(type) {
		case int:
			obj.Ptr = unsafe.Pointer(&val)
		default:
			obj.Encoding = EncodingRaw // 切换编码
			obj.Ptr = unsafe.Pointer(types.NewSDS(convert.String2Byte(val.(string))))
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
	case int:
		return NewStringObjectWithEncoding(val, EncodingInt)
	}

	return NewStringObjectWithEncoding(val, EncodingRaw)
}

func NewStringObjectWithEncoding(val any, encoding ObjectEncoding) *Object {
	obj := &Object{Type: ObjectString, Encoding: encoding}
	switch encoding {
	case EncodingInt:
		n, _ := val.(int)
		obj.Ptr = unsafe.Pointer(&n)
	default:
		obj.Ptr = unsafe.Pointer(types.NewSDS(convert.ToByte(val)))
	}
	return obj
}

func NewHashObject() *Object {
	obj := &Object{Type: ObjectHash, Encoding: EncodingHT, Ptr: unsafe.Pointer(types.NewHashTable())}

	return obj
}

func (obj *Object) SetHashField(field string, value any) error {
	if obj.Type != ObjectHash {
		return errors.InvalidType
	}

	table := (*types.HashTable)(obj.Ptr)
	table.Set(field, value)

	return nil
}

func (obj *Object) GetHashField(field string) any {
	if obj.Type != ObjectHash {
		return nil
	}

	table := (*types.HashTable)(obj.Ptr)
	return table.Get(field)
}

func (obj *Object) HasHashField(field string) bool {
	if obj.Type != ObjectHash {
		return false
	}

	table := (*types.HashTable)(obj.Ptr)
	return table.Has(field)

}

func (obj *Object) HLen() int {
	if obj.Type != ObjectHash {
		return 0
	}

	table := (*types.HashTable)(obj.Ptr)
	return table.Len()
}

func (obj *Object) HDel(fields ...string) (count int) {
	if obj.Type != ObjectHash {
		return 0
	}

	table := (*types.HashTable)(obj.Ptr)
	for _, field := range fields {
		if table.Has(field) {
			count++
			table.Del(field)
		}
	}
	return
}
