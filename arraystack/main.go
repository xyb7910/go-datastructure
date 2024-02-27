package arraystack

import "sync"

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

// Push 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.array = append(stack.array, v)

	stack.size++
}

// Pop 出栈
func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("stack is empty")
	}
	v := stack.array[stack.size-1]

	// 收缩切片，但空间可能会越拉越大
	//stack.array = stack.array[:stack.size-1]

	// 创建新的切片，长度为原来的长度-1，但是移动次数较多
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray

	stack.size--

	return v
}

// Peek 查看栈顶
func (stack *ArrayStack) Peek() string {
	if stack.size == 0 {
		panic("stack is empty")
	}
	v := stack.array[stack.size-1]
	return v
}

// Size 大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// IsEmpty 是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
