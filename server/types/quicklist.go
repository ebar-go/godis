package types

type QuickList struct {
	head, tail *QuickListNode
	count      uint64 // 所有压缩列表中的总元素个数
	len        uint64 // node节点个数
	fill       uint16 //zipList大小
	compress   uint16 //节点压缩深度
}

const (
	QuickListNodeEncodingRaw = 1 // 没有被压缩
	QuickListNodeEncodingLZF = 2 // 已经被LZF算法压缩
)

type QuickListNode struct {
	prev, next *QuickListNode // previous/next node
	zl         *ZipList
	size       uint64 // 压缩列表的字节大小
	count      uint64 // 压缩列表的元素个数
	encoding   uint   // 表示zipList是否被压缩
}

type ZipList struct{}

func (ql *QuickList) PushHead() {

}

func (ql *QuickList) PushTail()                 {}
func (ql *QuickList) InsertAfter()              {}
func (ql *QuickList) InsertBefore()             {}
func (ql *QuickList) ReplaceAtIndex(index uint) {}
func (ql *QuickList) DelEntry()                 {}
func (ql *QuickList) DelRange()                 {}
