package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc64"
)

func (store *Store) Set(key string, value string) {
	index := HashIndex(key)

	entry := store.dict.HashTable().Get(index)
	item := &DictEntry{
		Key: NewStringObject(key),
		Val: NewStringObject(value),
	}
	if entry == nil {
		store.dict.HashTable().Set(index, item)
		return
	}

	for {
		if entry.Key.String() == key {
			entry.Val = item.Val
			return
		}

		if entry.Next == nil {
			break
		}

		entry = entry.Next
	}

	entry.Next = item
	store.dict.HashTable().used++

}

func (store *Store) Get(key string) *Object {
	index := HashIndex(key)
	entry := store.dict.ht[0].table[index&store.dict.ht[0].mask]
	if entry == nil {
		return nil
	}

	for {
		if entry == nil {
			break
		}
		if entry.Key.String() == key {
			return entry.Val
		}

		entry = entry.Next

	}

	return nil
}

func (store *Store) Len(key string) uint64 {
	obj := store.Get(key)
	if obj == nil {
		return 0
	}

	return obj.Len()

}

func HashIndex(key any) uint64 {
	table := crc64.MakeTable(crc64.ECMA)
	bytesBuffer := bytes.NewBuffer([]byte{})
	switch key.(type) {
	case string:
		str := key.(string)
		err := binary.Write(bytesBuffer, binary.BigEndian, []byte(str))
		if err != nil {
			panic("invalid key")
		}
	default:
		err := binary.Write(bytesBuffer, binary.BigEndian, []byte(fmt.Sprintf("%v", key)))
		if err != nil {
			panic("invalid key")
		}
	}

	return crc64.Checksum(bytesBuffer.Bytes(), table)
}
