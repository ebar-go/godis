package server

import (
	"github.com/ebar-go/godis/server/types"
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
	EncodingSDS = iota
	EncodingQuickList
	EncodingListPack
	EncodingHashTable
	EncodingIntSet
	EncodingSkipList
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
		return uint64(sds.Len())
	}

	return 0
}

func (obj Object) String() string {
	switch obj.Type {
	case ObjectString:
		sds := (*types.SDS)(obj.Ptr)
		return sds.String()
	}

	return ""
}

func NewStringObject(str string) *Object {
	return &Object{
		Type:     ObjectString,
		Encoding: EncodingSDS,
		Ptr:      unsafe.Pointer(types.NewSDS(str)),
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
