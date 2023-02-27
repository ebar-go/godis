package types

type QuickList struct {
	head, tail *QuickListNode
	count      uint64 // 所有压缩列表中的总元素个数
	len        uint64 // node节点个数
	fill       uint16 //zipList大小
	compress   uint16 //节点压缩深度
}

func NewQuickList() *QuickList {
	head, tail := NewQuickListNode(), NewQuickListNode()
	head.next = tail
	tail.prev = head
	return &QuickList{
		head: head,
		tail: tail,
		len:  2,
	}
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

func NewQuickListNode() *QuickListNode {
	return &QuickListNode{
		zl: NewZipList(8),
	}
}

func (ql *QuickList) Len() uint64 {
	return ql.count
}

func (ql *QuickList) PushHead(value any) {
	entry := NewEntry(value)
	ql.count++
	if ql.head.zl.Full() {
		node := NewQuickListNode()
		node.next = ql.head
		node.zl.Insert(entry)
		ql.head = node
		ql.len++
		return
	}

	ql.head.zl.Insert(entry)
}

func (ql *QuickList) PushTail(value any) {
	entry := NewEntry(value)
	ql.count++
	if ql.tail.zl.Full() {
		node := NewQuickListNode()
		node.prev = ql.tail
		node.zl.Insert(entry)
		ql.tail = node
		ql.len++
		return
	}

	ql.tail.zl.Insert(entry)

}
func (ql *QuickList) InsertAfter()              {}
func (ql *QuickList) InsertBefore()             {}
func (ql *QuickList) ReplaceAtIndex(index uint) {}
func (ql *QuickList) DelEntry(entry *Entry) {
	node := ql.head
	for node != nil {
		if node.zl.Remove(entry) {
			ql.count--
			return
		}
		node = node.next
	}

}
func (ql *QuickList) DelRange() {

}

func (ql *QuickList) RPop(count int) []*Entry {
	items := make([]*Entry, 0, count)
	node := ql.head
	for node.zl.len == 0 {
		node = node.next
	}
	ql.count -= uint64(count)
	for count > 0 {
		diff := int(node.zl.len) - count
		if diff >= 0 {
			items = append(items, node.zl.entries[diff:node.zl.len]...)
			node.zl.entries = node.zl.entries[:diff]
			node.zl.len = uint8(diff)
			count = 0
		} else {
			diff = 0 - diff
			items = append(items, node.zl.entries[:node.zl.len]...)
			count = count - int(node.zl.len)
			node = node.next
		}

	}

	return items
}

func (ql *QuickList) LRange(start, end int64) []*Entry {
	if start > int64(ql.count) {
		return nil
	}
	if end > int64(ql.count) {
		end = int64(ql.count)
	}
	index := start
	node := ql.head
	n := end - start
	result := make([]*Entry, 0, n)
	// 找到应该从哪个node开始遍历
	for {
		if node == nil {
			return nil
		}
		zipListLen := int64(node.zl.len)
		if index < zipListLen {
			break
		}
		index -= zipListLen
		node = node.next
	}

	for {
		if n == 0 {
			break
		}

		zipListLen := int64(node.zl.len)
		if n <= zipListLen-index {
			result = append(result, node.zl.entries[index:index+n]...)
			break
		}

		result = append(result, node.zl.entries[index:]...)
		n -= zipListLen - index
		node = node.next
		index = 0

	}

	return result
}
