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

func NewSortedSetObject() *Object {
	return &Object{Type: ObjectSortedSet, Encoding: EncodingSkipList, Ptr: unsafe.Pointer(types.NewSkipList())}
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
func (obj *Object) RPush(values ...string) int {
	if obj.Type != ObjectList {
		return -1
	}

	list := (*types.QuickList)(obj.Ptr)
	count := len(values)
	for _, item := range values {
		list.PushHead(item)
	}
	return count

}

func (obj *Object) LLen() uint64 {
	if obj.Type != ObjectList {
		return 0
	}

	list := (*types.QuickList)(obj.Ptr)
	return list.Len()
}

func (obj *Object) LPop(count int) []string {
	if obj.Type != ObjectList {
		return nil
	}

	list := (*types.QuickList)(obj.Ptr)
	if uint64(count) > list.Len() {
		count = int(list.Len())
	}

	items := list.LRange(0, int64(count))
	res := make([]string, count)
	for idx, item := range items {
		res[idx] = item.Value.(string)
		list.DelEntry(item)
	}

	return res
}

func (obj *Object) RPop(count int) []string {
	if obj.Type != ObjectList {
		return nil
	}

	list := (*types.QuickList)(obj.Ptr)
	if uint64(count) > list.Len() {
		count = int(list.Len())
	}

	res := make([]string, count)
	items := list.RPop(count)
	for idx, item := range items {
		res[idx] = item.Value.(string)
	}

	return res
}

func (obj *Object) LRange(start, end int64) []string {
	if obj.Type != ObjectList {
		return nil
	}

	list := (*types.QuickList)(obj.Ptr)

	items := list.LRange(start, end)
	res := make([]string, len(items))
	for idx, item := range items {
		res[idx] = item.Value.(string)
	}

	return res
}

func (obj *Object) ZAdd(member string, score float64) int {
	if obj.Type != ObjectSortedSet {
		return 0
	}

	list := (*types.SkipList)(obj.Ptr)
	list.Insert(score, member)
	return 1
}

func (obj *Object) ZCard() int64 {
	if obj.Type != ObjectSortedSet {
		return 0
	}

	list := (*types.SkipList)(obj.Ptr)
	return list.Length()
}

func (obj *Object) ZRem(members ...string) int {
	if obj.Type != ObjectSortedSet {
		return 0
	}

	list := (*types.SkipList)(obj.Ptr)
	count := 0
	for _, member := range members {
		if list.Remove(member) {
			count++
		}
	}

	return count
}

func (obj *Object) ZScore(member string) (float64, bool) {
	if obj.Type != ObjectSortedSet {
		return 0, false
	}

	list := (*types.SkipList)(obj.Ptr)

	return list.Score(member)
}

func (obj *Object) ZRange(start, stop int64) []string {
	if obj.Type != ObjectSortedSet {
		return nil
	}

	list := (*types.SkipList)(obj.Ptr)
	nodes := list.Range(start, stop)
	items := make([]string, 0, len(nodes))
	for _, node := range nodes {
		items = append(items, node.Value.(string))
	}
	return items
}
