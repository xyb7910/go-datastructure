package tree

import (
	"fmt"
	"sync"
)

// TreeNode 定义树节点
type Node struct {
	Data  string // data
	Left  *Node  // left child
	Right *Node  // right child
}

// PreOrder 先序遍历 根左右
func PreOrder(tree *Node) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Data, " ")
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

// MidOrder 中序遍历 左根右
func MidOrder(tree *Node) {
	if tree == nil {
		return
	}

	MidOrder(tree.Left)
	fmt.Println(tree.Data, " ")
	MidOrder(tree.Right)
}

// PostOrder 后续遍历 左右根
func PostOrder(tree *Node) {
	if tree == nil {
		return
	}
	PostOrder(tree.Left)
	PostOrder(tree.Right)
	fmt.Println(tree.Data, " ")
}

// LayerOrder 层序遍历
func LayerOrder(tree *Node) {
	if tree == nil {
		return
	}

	// 借助队列实现层序遍历
	queue := new(LinkQueue)

	// 将根节点入队
	queue.Add(tree)

	// 层序遍历
	for queue.Size() > 0 {
		// 获取队列头元素
		element := queue.Remove()
		// 输出
		fmt.Println(element.Data, " ")

		// 将左右子树入队
		if element.Left != nil {
			queue.Add(element.Left)
		}

		if element.Right != nil {
			queue.Add(element.Right)
		}
	}
}

// LinkNode 定义链表节点
type LinkNode struct {
	Next  *LinkNode
	Value *Node
}

// LinkQueue 定义链表队列
type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// Add 入队
func (q *LinkQueue) Add(v *Node) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// 如果队列为空，我们将新节点作为队列的根节点
	if q.root == nil {
		q.root = new(LinkNode)
		q.root.Value = v
	} else {
		// 队列不为空，新建一个节点，采用尾插法实现
		newNode := new(LinkNode)
		newNode.Value = v

		// 找到尾节点
		nowNode := q.root
		if nowNode.Next != nil {
			nowNode = nowNode.Next
		}

		nowNode.Next = newNode
	}
	q.size++
}

// Remove 出队
func (q *LinkQueue) Remove() *Node {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		return nil
	}

	// 找到队头节点
	top := q.root
	v := top.Value

	// 将对头元素出队
	q.root = top.Next

	q.size--
	return v
}

// Size 队列大小
func (q *LinkQueue) Size() int {
	return q.size
}
