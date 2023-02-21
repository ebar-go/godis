package store

import (
	"github.com/ebar-go/godis/errors"
	"github.com/ebar-go/godis/internal/types"
	"github.com/ebar-go/godis/pkg/convert"
	"math/rand"
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

func NewSetObject() *Object {
	obj := &Object{Type: ObjectSet, Encoding: EncodingHT, Ptr: unsafe.Pointer(types.NewHashTable())}

	return obj
}

func NewListObject() *Object {
	return &Object{Type: ObjectList, Encoding: EncodingQuickList, Ptr: unsafe.Pointer(types.NewQuickList())}
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

func (obj *Object) HKeys() []string {
	if obj.Type != ObjectHash {
		return nil
	}

	table := (*types.HashTable)(obj.Ptr)
	return table.Fields()
}
func (obj *Object) HGetAll() map[string]any {
	if obj.Type != ObjectHash {
		return nil
	}

	table := (*types.HashTable)(obj.Ptr)
	return table.All()
}

func (obj *Object) SAdd(members ...string) error {
	if obj.Type != ObjectSet {
		return errors.InvalidType
	}

	table := (*types.HashTable)(obj.Ptr)
	for _, member := range members {
		table.Set(member, struct{}{})
	}

	return nil
}

func (obj *Object) SAddOrDie(members ...string) {
	if err := obj.SAdd(members...); err != nil {
		panic(err)
	}
}

func (obj *Object) SRem(members ...string) (count int, err error) {
	if obj.Type != ObjectSet {
		err = errors.InvalidType
		return
	}

	table := (*types.HashTable)(obj.Ptr)
	for _, member := range members {
		if table.Has(member) {
			count++
			table.Del(member)
		}
	}
	return
}

func (obj *Object) SCard() int64 {
	if obj.Type != ObjectSet {
		return 0
	}

	table := (*types.HashTable)(obj.Ptr)
	return int64(table.Len())
}

func (obj *Object) SPop(count int) []string {
	if obj.Type != ObjectSet {
		return nil
	}

	table := (*types.HashTable)(obj.Ptr)
	res := make([]string, 0, count)

	for i := 0; i < count; i++ {
		fields := table.Fields()
		idx := rand.Intn(len(fields))
		res = append(res, fields[idx])
		table.Del(fields[idx])
	}

	return res
}

func (obj *Object) SIsMember(member string) int {
	if obj.Type != ObjectSet {
		return 0
	}

	table := (*types.HashTable)(obj.Ptr)
	if table.Has(member) {
		return 1
	}
	return 0
}
func (obj *Object) SMembers() []string {
	if obj.Type != ObjectSet {
		return nil
	}

	table := (*types.HashTable)(obj.Ptr)
	return table.Fields()
}

func (obj *Object) LPush(values ...string) int {
	if obj.Type != ObjectList {
		return -1
	}

	list := (*types.QuickList)(obj.Ptr)
	count := len(values)
	for _, item := range values {
		list.PushTail(item)
	}
	return count

}
