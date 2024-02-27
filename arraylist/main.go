package arraylist

import (
	"fmt"
	"sync"
)

func ArrayLink() {
	type Value struct {
		Data      string
		NextIndex int64
	}

	var array [5]Value

	array[0] = Value{"a", 3}
	array[1] = Value{"b", 4}
	array[2] = Value{"c", 1}
	array[3] = Value{"d", 2}
	array[4] = Value{"e", -1}

	node := array[0]
	for {
		fmt.Println(node.Data)
		if node.NextIndex == -1 {
			break
		}
		node = array[node.NextIndex]
	}

}

// Array 定义可变长的数组
type Array struct {
	array []int       // 使用切片来代替
	len   int         //  表示数组的真实长度
	cap   int         // 容量
	lock  *sync.Mutex // 并发安全使用锁
}

func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len > cap")
	}

	// 把切片当做数组来用
	array := make([]int, cap, cap)

	// 元数据
	s.array = array
	s.len = 0
	s.cap = cap
	s.lock = &sync.Mutex{} // 初始化锁
	return s
}

// Append 向数组中添加一个元素，如果数组已满，则自动扩容
func (a *Array) Append(element int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == a.cap {
		newCap := a.len * 2

		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]int, newCap, newCap)

		// 把老的数据传输到新的数组里边
		for k, v := range a.array {
			newArray[k] = v
		}

		a.array = newArray
		a.cap = newCap
	}

	a.array[a.len] = element
	a.len = a.len + 1
}

// AppendMany 向数组中添加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

// Get 通过下标获取元素
func (a *Array) Get(index int) int {
	// 处理越界
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.array[index]
}

// Len 返回真实的长度
func (a *Array) Len() int {
	return a.len
}

// Cap 返回真实的容量
func (a *Array) Cap() int {
	return a.cap
}

func PrintArray(array *Array) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 获取第一个元素
		if i == 0 {
			result += fmt.Sprintf("%s%d", result, array.Get(i))
			continue
		}

		result = fmt.Sprintf("%s, %d", result, array.Get(i))
	}
	result = result + "]"
	return
}
