package types

type ListPack struct {
	size    uint
	entries []ListPackEntry
}

type ListPackEntry struct {
	Value any
}

func NewListPack(capacity uint) *ListPack {
	return &ListPack{entries: make([]ListPackEntry, 0, capacity)}
}

func (lp *ListPack) Insert(val any) {
	entry := ListPackEntry{Value: val}
	lp.entries = append(lp.entries, entry)
	lp.size++
}
