package main

import "fmt"

// 单链表
type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	node := new(LinkNode)

	node.Data = 1

	node1 := new(LinkNode)
	node1.Data = 2
	node.NextNode = node1

	node2 := new(LinkNode)
	node2.Data = 3
	node1.NextNode = node2

	// print list data
	nowNode := node
	for {
		if nowNode != nil {
			fmt.Println(nowNode.Data)

			nowNode = nowNode.NextNode
			continue
		}
		break
	}

}
