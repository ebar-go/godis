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

	newNode := &SkipListNode{score: score, Value: val, levels: make([]SkipListLevel, level)}
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

func getRandLevel() uint {
	level := uint(1)
	for rand.Float32() < p && level < SkipListMaxLevel {
		level++
	}
	return level
}
