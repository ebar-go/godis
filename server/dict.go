package server

const (
	defaultMask = 1 << 5
)

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
