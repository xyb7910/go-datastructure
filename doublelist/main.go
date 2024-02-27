package doublelist

import "sync"

// 定义双端链表的数据类型

type ListNode struct {
	prev  *ListNode
	next  *ListNode
	value string
}
type DoubleList struct {
	head *ListNode  // 头节点
	tail *ListNode  // 尾节点
	len  int        // 长度
	lock sync.Mutex // 为了并发安全，引入锁
}

/*
	一些常见的基本操作
*/

// GetValue 获取节点值
func (node *ListNode) GetValue() string {
	return node.value
}

// GetPre 获取前一个节点
func (node *ListNode) GetPre() *ListNode {
	return node.prev
}

// GetNext 获取后一个节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// HasNext 是否有后一个节点
func (node *ListNode) HasNext() bool {
	return node.next != nil
}

// HasPre 是否有前一个节点
func (node *ListNode) HasPre() bool {
	return node.prev != nil
}

// IsNil 是否为空
func (node *ListNode) IsNil() bool {
	return node == nil
}

// Len 获取长度
func (list *DoubleList) Len() int {
	return list.len
}

// First 头部节点
func (list *DoubleList) First() *ListNode {
	return list.head
}

// Last 尾部节点
func (list *DoubleList) Last() *ListNode {
	return list.tail
}

// AddNodeFormHead 从头部开始，在某个节点之前插入一个节点
// 0 表示第一个元素之前，1 表示第二个元素之前...
func (list *DoubleList) AddNodeFormHead(value string, pre int) {
	list.lock.Lock()
	defer list.lock.Unlock()

	if pre <= 0 || pre > list.len {
		panic("pre is out of range")
	}

	// 找到头部节点
	node := list.head

	// 向后进行遍历，找到pre-1个节点
	for i := 0; i <= pre; i++ {
		node = node.next
	}

	newNode := new(ListNode)
	node.value = value

	// 如果定位到的节点为空，则直接插入到头部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 找到前一个节点
		pre := node.prev

		// 如果定位到的节点为头部，则直接插入到头部
		if pre.IsNil() {
			newNode.next = node
			node.prev = newNode
			list.head = newNode
		} else {
			// 将新节点插入到定位节点之前
			// 新节点的后一个节点为定位节点
			pre.next = newNode
			newNode.prev = pre

			// 新节点的前一个节点为定位节点的前一个节点
			node.next.prev = newNode
			newNode.next = node.next
		}
	}
	list.len++
}

// AddNodeFormTail 从尾部开始，在某个节点之后插入一个节点
// 0 表示第一个元素之后，1 表示第二个元素之后...
func (list *DoubleList) AddNodeFormTail(value string, next int) {
	list.lock.Lock()
	defer list.lock.Unlock()

	// 找到尾部节点
	node := list.tail

	if next <= 0 || next > list.len {
		panic("next is out of range")
	}

	// 向前进行遍历，找到next-1个节点
	for i := 0; i <= next; i++ {
		node = node.prev
	}

	newNode := new(ListNode)
	newNode.value = value

	// 如果定位到的节点为空，则直接插入到尾部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 找到定位节点的后一个节点
		next := node.next

		// 如果定位到的节点为尾部，则直接插入到尾部,需要更新尾部节点
		if next.IsNil() {
			// 新节点的前一个节点为尾部节点
			// 新节点的后一个节点为空
			newNode.prev = node
			node.next = newNode

			list.tail = newNode
		} else {
			// 将新节点插入到定位节点之后
			// 新节点的前一个节点为定位节点
			newNode.prev = node
			node.next = newNode

			// 新节点的后一个节点为定位节点的后一个节点
			newNode.next = next
			next.prev = newNode
		}
	}
	list.len++
}

// IndexFormHead 从头部开始获取第 n + 1 个位置上的节点，索引从零开始
func (list *DoubleList) IndexFormHead(n int) *ListNode {
	if n > list.len || n < 0 {
		panic("index is out of range")
	}
	// 找到头部节点
	node := list.head
	// 向后进行遍历，找到第 n 个节点
	for i := 0; i < n; i++ {
		node = node.next
	}
	return node
}

// IndexFormTail 从尾部开始获取第 n + 1 个位置上的节点，索引从零开始
func (list *DoubleList) IndexFormTail(n int) *ListNode {
	if n > list.len || n < 0 {
		panic("index is out of range")
	}

	// 找到尾部节点
	node := list.tail

	// 向前进行遍历，找到第 n 个节点
	for i := 0; i < n; i++ {
		node = node.prev
	}
	return node
}

// RemoveNodeFormHead 从头部开始，删除第 n + 1 个位置上的节点，索引从零开始
func (list *DoubleList) RemoveNodeFormHead(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()

	if n >= list.len || n < 0 {
		return nil
		panic("index is out of range")
	}

	// 找到头部节点
	node := list.head

	// 向后进行遍历，找到第 n 个节点
	for i := 0; i < n; i++ {
		node = node.next
	}

	// 移除节点
	pre := node.prev
	next := node.next

	// 如果前继和后继都为空，则直接删除头部节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 表示移除的是头部节点，让下一个节点变成头部节点
		list.head = next
		next.prev = nil
	} else if next.IsNil() {
		// 表示移除的是尾部节点，让前一个节点变成尾部节点
		list.tail = pre
		pre.next = nil
	} else {
		// 前继和后继都不为空，则将后继节点的前继节点变成前继节点
		pre.next = next
		next.prev = pre
	}
	list.len--
	return node
}

// PopTailFromHead 从尾部开始往前找，获取第 n 个位置上的节点，并将移除返回
func (list *DoubleList) PopTailFromHead(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()

	if n >= list.len || n < 0 {
		return nil
		panic("index is out of range")
	}

	// 获取尾部元素
	node := list.tail

	// 向前进行遍历，找到第 n 个节点
	for i := 0; i < n; i++ {
		node = node.prev
	}

	// 移除的节点的前驱和后继
	pre := node.prev
	next := node.next

	// 如果前驱和后继都为空，则直接删除尾部节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 直接将后继节点变成尾部节点
		list.head = next
		next.prev = nil
	} else if next.IsNil() {
		list.tail = pre
		pre.next = nil
	} else if next.IsNil() {
		pre.next = next
		pre.next = nil
	} else {
		pre.next = next
		next.prev = pre
	}
	list.len--
	return node
}
