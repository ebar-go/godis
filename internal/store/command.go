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

func (store *Store) SRem(key string, members ...string) (int, error) {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.SRem(members...)
		}

		entry = entry.Next
	}

	return 0, nil
}

func (store *Store) SCard(key string) int64 {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.SCard()
		}

		entry = entry.Next
	}

	return 0
}

func (store *Store) SPop(key string, count int) []string {
	if count <= 0 {
		count = 1
	}
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.SPop(count)
		}

		entry = entry.Next
	}

	return nil
}

func (store *Store) SIsMember(key string, member string) int {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.SIsMember(member)
		}

		entry = entry.Next
	}

	return 0
}

func (store *Store) SMembers(key string) []string {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.SMembers()
		}

		entry = entry.Next
	}

	return nil
}

func (store *Store) LPush(key string, val ...string) int {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.LPush(val...)
		}

		entry = entry.Next
	}

	obj := NewListObject()
	count := obj.LPush(val...)
	store.dict.HashTable().Set(index, &DictEntry{
		Key: NewKeyObject(key),
		Val: obj,
	})

	store.dict.HashTable().Increment()

	return count
}

func (store *Store) RPush(key string, val ...string) int {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.RPush(val...)
		}

		entry = entry.Next
	}

	obj := NewListObject()
	count := obj.RPush(val...)
	store.dict.HashTable().Set(index, &DictEntry{
		Key: NewKeyObject(key),
		Val: obj,
	})

	store.dict.HashTable().Increment()

	return count
}

func (store *Store) LLen(key string) uint64 {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.LLen()
		}

		entry = entry.Next
	}

	return 0
}

func (store *Store) LPop(key string, count int) []string {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.LPop(count)
		}

		entry = entry.Next
	}

	return nil
}

func (store *Store) RPop(key string, count int) []string {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.RPop(count)
		}

		entry = entry.Next
	}

	return nil
}

func (store *Store) LRange(key string, start, end int64) []string {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.LRange(start, end)
		}

		entry = entry.Next
	}

	return nil
}

func (store *Store) ZAdd(key string, member string, score float64) int {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.ZAdd(member, score)
		}

		entry = entry.Next
	}

	obj := NewSortedSetObject()
	count := obj.ZAdd(member, score)
	store.dict.HashTable().Set(index, &DictEntry{
		Key: NewKeyObject(key),
		Val: obj,
	})

	store.dict.HashTable().Increment()

	return count
}

func (store *Store) ZCard(key string) int64 {
	index := HashIndex(key)
	entry := store.dict.HashTable().Get(index)

	for {
		if entry == nil {
			break
		}

		if entry.Key.String() == key {
			return entry.Val.ZCard()
		}

		entry = entry.Next
	}

	return 0
}
