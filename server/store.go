package server

import (
	"github.com/ebar-go/godis/pkg/convert"
	"github.com/ebar-go/godis/server/types"
	"strconv"
	"unsafe"
)

type Store struct {
	dict *Dict
}

func NewStore() *Store {
	return &Store{dict: NewDict()}
}

type Dict struct {
	ht [2]*DictHT
}

func (dict *Dict) HashTable() *DictHT {
	return dict.ht[0]
}

func (dict *Dict) SetEntry(index uint64, key string, value any) {
	ht := dict.HashTable()
	entry := ht.Get(index)
	if entry == nil {
		ht.Set(index, &DictEntry{
			Key: NewStringObject([]byte(key)),
			Val: NewObject(value),
		})
		return
	}

	// 通过链表寻址解决hash冲突
	for {
		if entry.Key.String() == key {
			entry.Val.SetStringValue(value)
			return
		}

		if entry.Next == nil {
			break
		}

		entry = entry.Next
	}

	entry.Next = &DictEntry{
		Key: NewStringObject([]byte(key)),
		Val: NewObject(value),
	}
	ht.used++
}

func NewDict() *Dict {
	return &Dict{ht: [2]*DictHT{
		{
			table: make([]*DictEntry, 128),
			mask:  31,
			size:  128,
			used:  0,
		},
	}}
}

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

func NewObject(val any) *Object {
	switch val.(type) {
	case string:
		return NewStringObject(convert.ToByte(val))
	case int:
		return NewIntObject(val.(int))
	}

	return nil
}

func NewStringObject(buf []byte) *Object {
	return &Object{
		Type:     ObjectString,
		Encoding: EncodingRaw,
		Ptr:      unsafe.Pointer(types.NewSDS(buf)),
	}
}

func NewIntObject(n int) *Object {
	return &Object{
		Type:     ObjectString,
		Encoding: EncodingInt,
		Ptr:      unsafe.Pointer(&n),
	}
}

type DictHT struct {
	table []*DictEntry
	mask  uint64
	size  uint64
	used  uint64
}

func (ht *DictHT) Get(index uint64) *DictEntry {
	return ht.table[index&ht.mask]
}

func (ht *DictHT) Set(index uint64, entry *DictEntry) {
	ht.used++
	ht.table[index&ht.mask] = entry
}

type DictEntry struct {
	Key  *Object
	Val  *Object
	Next *DictEntry
}
