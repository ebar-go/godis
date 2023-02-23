package store

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

func (ht *DictHT) Increment() {
	ht.used++
}
