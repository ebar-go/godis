package types

var (
	ZipListEnd = []byte{255}
)

const (
	ZipListHeaderLen = 80
)

// ZipList 压缩列表，redis3.0之前用于实现list/hash/zset,3.0以后被QuickList替代
type ZipList struct {
	bytes   uint32 // 总长度(单位: 字节)
	tail    uint32 // 最后一个节点的偏移量
	len     uint16 // 元素个数
	entries []byte
}

func NewZipList(n uint32) *ZipList {
	return &ZipList{
		bytes:   ZipListHeaderLen + n, // 32(bytes) + 32(tail) + 16(len) + n
		entries: make([]byte, n),
	}
}

// Entry 节点元素,由PrevLength,Encoding,Data三部分组成
// entry的前8位小于254，则这8位就表示上一个节点的长度
// entry的前8位等于254，则意味着上一个节点的长度无法用8位表示，后面32位才是真实的prevLength
// 用254 不用255(11111111)作为分界是因为255是zlEnd的值，它用于判断zipList是否到达尾部。
// 故entry的prev长度会是1或者5字节
type Entry []byte

const (
	EntryEncodingInt16      = (0xc0 | 0<<4)
	EntryEncodingInt32      = (0xc0 | 1<<4)
	EntryEncodingInt64      = (0xc0 | 2<<4)
	EntryEncodingInt24      = (0xc0 | 3<<4)
	EntryEncodingInt8       = 0xfe
	EntryEncodingIntIMMMask = 0x0f
	EntryEncodingIntIMMMin  = 0xf1
	EntryEncodingIntIMMMax  = 0xfd
	EntryEncodingStr6       = 0 << 6
	EntryEncodingStr14      = 1 << 6
	EntryEncodingStr32      = 2 << 6
)

func (zl *ZipList) Insert(entry Entry) {
	if zl.len == 0 {
		copy(zl.entries, entry)
		copy(zl.entries[len(entry):], ZipListEnd)
		zl.tail = ZipListHeaderLen + uint32(len(entry))
		return
	}

}
