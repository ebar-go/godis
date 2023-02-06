package types

var (
	ZipListEnd = []byte{255}
)

const (
	ZipListHeaderLen = 80
)

// ZipList 压缩列表，redis3.0之前用于实现list/hash/zset,3.0以后被QuickList替代
type ZipList struct {
	len     uint8 // 元素个数
	entries []*Entry
}

func NewZipList(n uint8) *ZipList {
	return &ZipList{
		entries: make([]*Entry, n),
	}
}

type Entry struct {
	Value any
}

func NewEntry(value any) *Entry {
	return &Entry{Value: value}
}

func (zl *ZipList) Full() bool {
	return zl.len == uint8(len(zl.entries))
}

func (zl *ZipList) Insert(entry *Entry) {
	zl.entries[zl.len] = entry
	zl.len++
}
