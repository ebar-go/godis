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

func NewDict() *Dict {
	return &Dict{ht: [2]*DictHT{}}
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

func NewStringObject(str string) *Object {
	return &Object{
		Type:     ObjectString,
		Encoding: EncodingSDS,
		Ptr:      unsafe.Pointer(types.NewSDS(str)),
	}
}

type DictHT struct {
	entries []DictEntry
}

type DictEntry struct {
	Key *Object
	Val *Object
}
