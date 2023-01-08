package server

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
	return &Dict{ht: [2]*DictHT{
		{
			table: make([]*DictEntry, 128),
			mask:  31,
			size:  128,
			used:  0,
		},
	}}
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
