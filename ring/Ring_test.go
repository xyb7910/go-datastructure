package main

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	r := &Ring{Value: 1}

	// link five nodes
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	r.Link(&Ring{Value: 5})

	node := r
	for {
		// print node value
		fmt.Println(node.Value)
		// move to next node
		node = node.Next()
		// if node reaches the end, go back to the beginning and end
		if node == r {
			return
		}
	}
}

func TestRing_Unlink(t *testing.T) {
	// 第一个节点
	r := &Ring{Value: 1}

	// 链接新的五个节点
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	r.Link(&Ring{Value: 5})

	temp := r.Unlink(3)

	node := r

	for {
		fmt.Println(node.Value)
		node = node.Next()
		if node == r {
			break
		}
	}

	fmt.Println("--------")

	node = temp
	for {
		fmt.Println(node.Value)
		node = node.Next()

		if node == temp {
			break
		}
	}
}
