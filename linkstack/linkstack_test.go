package linkstack

import (
	"fmt"
	"testing"
)

func TestLinkStack(t *testing.T) {
	linkstack := new(LinkStack)
	linkstack.Push("cat")
	linkstack.Push("dog")
	linkstack.Push("hen")
	fmt.Println("size", linkstack.Size())
	fmt.Println("peek", linkstack.Peek())
	fmt.Println("pop", linkstack.Pop())
	fmt.Println("size", linkstack.Size())
	fmt.Println("is empty", linkstack.IsEmpty())
}
