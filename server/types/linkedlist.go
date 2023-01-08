package types

import "unsafe"

// LinkedList Redis 的 List 对象的底层实现之一就是链表
type LinkedList struct {
	head, tail *ListNode // 首/尾节点

	len uint64 // 链表长度
}

// NewLinkedList returns a new linked list instance
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Len 返回list长度,时间复杂度:O(1)
func (ll *LinkedList) Len() uint64 {
	return ll.len
}

// First 返回list的头部节点，时间复杂度:O(1)
func (ll *LinkedList) First() *ListNode {
	return ll.head
}

// Last 返回list的尾部节点，时间复杂度:O(1)
func (ll *LinkedList) Last() *ListNode {
	return ll.tail
}

func (ll *LinkedList) Release() {
	node := ll.head
	for {
		if node == nil {
			break
		}
		node.prev = nil
		node = node.next
	}
	ll.head = nil
	ll.tail = nil
	ll.len = 0
}

// AddNodeHead 将一个节点添加到链表的表头
func (ll *LinkedList) AddNodeHead(value unsafe.Pointer) {
	node := &ListNode{value: value}
	if ll.len == 0 {
		ll.head = node
		ll.tail = node
		ll.len++
		return
	}

	node.next = ll.head
	ll.head.prev = node
	ll.head = node
	ll.len++
}

// AddNodeTail 将一个节点添加到链表的表尾
func (ll *LinkedList) AddNodeTail(value unsafe.Pointer) {
	node := &ListNode{value: value}
	if ll.len == 0 {
		ll.head = node
		ll.tail = node
		ll.len++
		return
	}

	node.prev = ll.tail
	ll.tail.next = node
	ll.tail = node
	ll.len++
}

func (ll *LinkedList) InsertNode(old *ListNode, value unsafe.Pointer, after bool) {
	node := &ListNode{value: value}
	if after {
		node.prev = old
		node.next = old.next
		if ll.tail == old {
			ll.tail = node
		}
	} else {
		node.next = old
		node.prev = old.prev
		if ll.head == old {
			ll.head = node
		}
	}

	if node.prev != nil {
		node.prev.next = node
	}

	if node.next != nil {
		node.next.prev = node
	}
	ll.len++
}

func (ll *LinkedList) DelNode(node *ListNode) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		ll.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		ll.tail = node.prev
	}
	ll.len--
}

// dup 复制节点值
func (ll *LinkedList) dup() {}

// free 释放节点
func (ll *LinkedList) free() {}

// match 比较节点值
func (ll *LinkedList) match() {}

// ListNode 链表节点
type ListNode struct {
	prev  *ListNode
	next  *ListNode
	value unsafe.Pointer
}
