package store

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc64"
)

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

func (store *Store) Set(key string, value any) {
	index := HashIndex(key)
	store.dict.SetEntry(index, key, value)
}

func (store *Store) Get(key string) *Object {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

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

func (store *Store) Del(key string) bool {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)
	if entry == nil {
		return false
	}

	if entry.Key.String() == key {
		store.dict.HashTable().Set(index, entry.Next)
		store.dict.HashTable().used--
		return true
	}

	for {
		if entry.Next == nil {
			break
		}

		if entry.Next.Key.String() == key {
			entry.Next = entry.Next.Next
			store.dict.HashTable().used--
			return true
		}

		entry = entry.Next
	}

	return false

}

func (store *Store) Has(key string) bool {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return true
		}

		entry = entry.Next

	}

	return false
}

func (store *Store) SetExpire(key string, ttl int64) {
	index := HashIndex(key)
	store.dict.SetExpire(index, key, ttl)
}

func (store *Store) GetExpire(key string) int64 {
	index := HashIndex(key)
	return store.dict.GetExpire(index, key)
}

func (store *Store) HSet(key string, field string, value any) error {
	index := HashIndex(key)

	return store.dict.SetHash(index, key, field, value)
}

func (store *Store) HGet(key string, field string) any {
	index := HashIndex(key)
	return store.dict.HGet(index, key, field)
}

func (store *Store) HExists(key string, field string) bool {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.HasHashField(field)
		}

		entry = entry.Next
	}

	return false
}

func (store *Store) HLen(key string) int64 {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return int64(entry.Val.HLen())
		}

		entry = entry.Next
	}

	return 0
}

func (store *Store) HDel(key string, fields ...string) int {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.HDel(fields...)
		}

		entry = entry.Next
	}

	return 0
}

func (store *Store) HKeys(key string) []string {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.HKeys()
		}

		entry = entry.Next
	}

	return nil
}
func (store *Store) HGetAll(key string) map[string]any {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.HGetAll()
		}

		entry = entry.Next
	}

	return nil
}

func (store *Store) SAdd(key string, members ...string) error {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.SAdd(members...)
		}

		entry = entry.Next
	}

	obj := NewSetObject()
	obj.SAddOrDie(members...)
	store.dict.HashTable().Set(index, &DictEntry{
		Key: NewKeyObject(key),
		Val: obj,
	})

	store.dict.HashTable().used++

	return nil
}
