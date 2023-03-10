package types

import "math/rand"

const (
	SkipListMaxLevel = 1 << 6
	p                = 0.25
)

type SkipList struct {
	head, tail *SkipListNode
	length     int64
	level      uint
}

func NewSkipList() *SkipList {
	sl := &SkipList{level: 1}
	sl.head = &SkipListNode{levels: make([]SkipListLevel, SkipListMaxLevel)}
	sl.tail = nil
	return sl
}

type SkipListNode struct {
	Value    any
	score    float64
	backward *SkipListNode
	levels   []SkipListLevel
}

type SkipListLevel struct {
	forward *SkipListNode
	span    uint
}

func (sl *SkipList) Insert(score float64, val any) {
	update := make(map[uint]*SkipListNode, SkipListMaxLevel)
	node := sl.head

	for i := int(sl.level) - 1; i >= 0; i-- {
		for node.levels[i].forward != nil && node.levels[i].forward.score < score {
			node = node.levels[i].forward
		}

		update[uint(i)] = node
	}
	node = node.levels[0].forward
	if node != nil && node.score == score {
		return
	}

	level := getRandLevel()

	newNode := &SkipListNode{score: score, Value: val, levels: make([]SkipListLevel, level), backward: node}
	for i := uint(0); i < level; i++ {
		newNode.levels[i].forward = update[i].levels[i].forward
		update[i].levels[i].forward = newNode
	}
	sl.length++

}

func (sl *SkipList) Search(score float64) (*SkipListNode, bool) {
	node := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for node.levels[i].forward != nil && node.levels[i].forward.score < score {
			node = node.levels[i].forward
		}
	}

	node = node.levels[0].forward
	if node != nil && node.score == score {
		return node, true
	}

	return nil, false
}

func (sl *SkipList) Length() int64 {
	return sl.length
}

func (sl *SkipList) Score(member any) (float64, bool) {
	node := sl.head
	for node != nil {
		if node.Value == member {
			return node.score, true
		}
		node = node.levels[0].forward
	}

	return 0, false
}

func (sl *SkipList) Remove(member any) bool {
	node := sl.head
	for node != nil {
		if node.Value == member {
			for i := 0; i < len(node.levels); i++ {
				if node.backward == nil {

				} else {
					node.backward.levels[i].forward = node.levels[i].forward
				}

			}
			node.levels[0].forward.backward = node.backward
			sl.length--
			return true
		}
		node = node.levels[0].forward
	}

	return false
}

func (sl *SkipList) Range(start, stop int64) []*SkipListNode {
	if stop > sl.length {
		stop = sl.length
	}
	nodes := make([]*SkipListNode, 0, stop-start+1)

	node := sl.head
	for i := int64(0); i < start; i++ {
		node = node.levels[0].forward
	}

	for idx := start; idx < stop; idx++ {
		node = node.levels[0].forward
		nodes = append(nodes, node)
	}

	return nodes
}

func getRandLevel() uint {
	level := uint(1)
	for rand.Float32() < p && level < SkipListMaxLevel {
		level++
	}
	return level
}
