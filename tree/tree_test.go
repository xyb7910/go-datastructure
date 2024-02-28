package tree

import (
	"fmt"
	"testing"
)

func TestOrder(f *testing.T) {
	t := &Node{Data: "A"}
	t.Right = &Node{Data: "B"}
	t.Left = &Node{Data: "C"}
	t.Left.Left = &Node{Data: "D"}
	t.Left.Right = &Node{Data: "E"}
	t.Right.Right = &Node{Data: "F"}

	fmt.Println("先序遍历：")
	PreOrder(t)
	fmt.Println("中序遍历：")
	MidOrder(t)
	fmt.Println("后续遍历：")
	PostOrder(t)
}

func TestLayerOrder(f *testing.T) {
	t := &Node{Data: "A"}
	t.Right = &Node{Data: "B"}
	t.Left = &Node{Data: "C"}
	t.Left.Left = &Node{Data: "D"}
	t.Left.Right = &Node{Data: "E"}
	t.Right.Right = &Node{Data: "F"}
	fmt.Println("层序遍历：")
	LayerOrder(t)
}
