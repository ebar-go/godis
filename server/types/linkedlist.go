package types

import "unsafe"

// LinkedList Redis 的 List 对象的底层实现之一就是链表
type LinkedList struct {
	head, tail *ListNode // 首/尾节点

	len uint64 // 链表长度
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
