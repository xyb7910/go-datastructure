package linkstack

import "sync"

type LinkNode struct {
	Value string
	Next  *LinkNode
}

type LinkStack struct {
	root *LinkNode // 链表起点
	size int
	lock sync.Mutex
}

// Push 入栈
func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 表示栈目前为空，直接向栈顶插入元素
	if stack.root == nil {
		stack.root = &LinkNode{Value: v, Next: nil}
	} else {
		// 将元素插入到栈顶 , 头插法
		preNode := stack.root

		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v

		newNode.Next = preNode

		// 更新 root
		stack.root = newNode
	}
	stack.size++
}

// Pop 弹出栈顶元素
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("stack is empty")
	}

	// 栈顶元素
	topNode := stack.root
	stack.root = topNode.Next

	stack.size--

	return topNode.Value
}

// Peek 栈顶元素
func (stack *LinkStack) Peek() string {
	// 栈为空
	if stack.size == 0 {
		panic("stack is empty")
	}
	// 栈顶元素
	v := stack.root.Value
	return v
}

// Size 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}
