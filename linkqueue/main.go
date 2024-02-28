package linkqueue

import "sync"

type LinkNode struct {
	Next  *LinkNode
	Value string
}

type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// Size 获取队列长度
func (queue *LinkQueue) Size() int {
	return queue.size
}

// Add 新元素入队
func (queue *LinkQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队列为空直接入队, 尾查法
	if queue.root == nil {
		queue.root = &LinkNode{Value: v, Next: nil}
	} else {
		newNode := new(LinkNode)
		newNode.Value = v

		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		nowNode.Next = newNode

		queue.size++
	}
}

func (queue *LinkQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队列为空直接返回
	if queue.root == nil {
		panic("queue is empty")
	}
	// 头部出元素
	topNode := queue.root
	v := topNode.Value

	queue.root = queue.root.Next
	queue.size--
	return v
}
