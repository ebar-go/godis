package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc64"
)

func (store *Store) Set(key string, value any) {
	index := HashIndex(key)
	store.dict.SetEntry(index, key, value)
}

func (store *Store) Get(key string) *Object {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)
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

func (store *Store) Del(key string) {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)
	if entry == nil {
		return
	}

	if entry.Key.String() == key {
		store.dict.HashTable().Set(index, entry.Next)
		store.dict.HashTable().used--
		return
	}

	for {
		if entry.Next == nil {
			return
		}

		if entry.Next.Key.String() == key {
			entry.Next = entry.Next.Next
			store.dict.HashTable().used--
			return
		}

		entry = entry.Next
	}

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
