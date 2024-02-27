package arraystack

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	arraySatck := new(ArrayStack)
	arraySatck.Push("cat")
	arraySatck.Push("dog")
	arraySatck.Push("hen")
	fmt.Println("size", arraySatck.Size())
	fmt.Println("peek", arraySatck.Peek())
	fmt.Println("pop", arraySatck.Pop())
	fmt.Println("size", arraySatck.Size())
	fmt.Println("is empty", arraySatck.IsEmpty())
}
