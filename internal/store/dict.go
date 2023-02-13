package store

import (
	"github.com/ebar-go/godis/constant"
	"strconv"
	"time"
)

const (
	defaultMask = 1 << 5
)

type Dict struct {
	ht [2]*DictHT
}

func (dict *Dict) HashTable() *DictHT {
	return dict.ht[0]
}

func (dict *Dict) SetHash(index uint64, key string, field string, value any) error {
	ht := dict.HashTable()
	entry := ht.Get(index)

	for {
		if entry == nil {
			obj := NewHashObject()
			_ = obj.SetHashField(field, value)
			ht.Set(index, &DictEntry{
				Key: NewKeyObject(key),
				Val: obj,
			})

			ht.used++
			break
		}

		if entry.Key.String() == key {
			return entry.Val.SetHashField(field, value)
		}

		entry = entry.Next
	}

	return nil
}

func (dict *Dict) ExpireTable() *DictHT {
	return dict.ht[1]
}

func (dict *Dict) SetExpire(index uint64, key string, ttl int64) {
	ht := dict.ExpireTable()
	entry := ht.Get(index)
	var expired int64
	if ttl > 0 {
		expired = time.Now().Add(time.Second * time.Duration(ttl)).Unix()
	}

	if entry == nil {
		ht.Set(index, &DictEntry{
			Key: NewKeyObject(key),
			Val: NewStringObject(strconv.FormatInt(expired, 10)),
		})
		return
	}

	// 通过链表寻址解决hash冲突
	for {
		if entry.Key.String() == key {
			entry.Val.SetStringValue(expired)
			return
		}

		if entry.Next == nil {
			break
		}

		entry = entry.Next
	}

	entry.Next = &DictEntry{
		Key: NewKeyObject(key),
		Val: NewStringObject(expired),
	}

}

func (dict *Dict) GetExpire(index uint64, key string) int64 {
	ht := dict.ExpireTable()
	entry := ht.Get(index)
	if entry == nil {
		// key exists and not set expiration
		return constant.ExpireResultOfForever
	}

	for {
		if entry.Key.String() == key {
			expired, _ := strconv.ParseInt(entry.Val.String(), 10, 64)
			now := time.Now().Unix()
			if expired <= now {
				return constant.ExpireResultOfExpired
			}

			return expired - now
		}

		if entry.Next == nil {
			break
		}

		entry = entry.Next
	}

	// not found
	return constant.ExpireResultOfForever
}
func (dict *Dict) SetEntry(index uint64, key string, value any) {
	ht := dict.HashTable()
	entry := ht.Get(index)
	if entry == nil {
		ht.Set(index, &DictEntry{
			Key: NewKeyObject(key),
			Val: NewStringObject(value),
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
		Key: NewKeyObject(key),
		Val: NewStringObject(value),
	}
	ht.used++
}

func NewDict() *Dict {
	return NewDictWithMask(defaultMask)
}

func NewDictWithMask(n int) *Dict {
	n = roundUp(n)
	return &Dict{ht: [2]*DictHT{
		{
			table: make([]*DictEntry, n),
			mask:  uint64(n - 1),
			size:  uint64(n),
			used:  0,
		},
		{
			table: make([]*DictEntry, n),
			mask:  uint64(n - 1),
			size:  uint64(n),
			used:  0,
		},
	}}
}

// roundUp 返回邻近的2的N次方的数
func roundUp(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}
