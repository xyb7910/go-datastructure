package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{3, 2, 1, 5, 4}
	BubbleSort(list)
	fmt.Println(list)
}

func TestSelectSortSort(t *testing.T) {
	list := []int{3, 2, 1, 5, 4}
	SelectSort(list)
	fmt.Println(list)
}

func TestSelectProSortSort(t *testing.T) {
	list := []int{3, 2, 1, 5, 4}
	SelectProSort(list)
	fmt.Println(list)
}

func TestInsertSortSort(t *testing.T) {
	list := []int{3, 2, 1, 5, 4}
	InsertSort(list)
	fmt.Println(list)
}
