package doublelist

import (
	"fmt"
	"testing"
)

func TestDoubleList(t *testing.T) {
	list := new(DoubleList)

	// 在列表头添加元素
	list.AddNodeFormHead("I", 0)
	list.AddNodeFormHead("love", 0)
	list.AddNodeFormHead("you", 0)

	// 在列表尾添加元素
	list.AddNodeFormTail("may", 0)
	list.AddNodeFormTail("happy", 0)

	list.AddNodeFormTail("begin second", list.Len()-1)
	list.AddNodeFormHead("begin second", list.Len()-1)

	// 比较慢的，因为内部会遍历拿到的值

	for i := 0; i < list.Len(); i++ {
		node := list.IndexFormHead(i)

		if !node.IsNil() {
			fmt.Println(node.GetValue())
		}
	}

}
