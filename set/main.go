package set

import (
	"sync"
)

// 思想：不考虑字典的值，我们可以实现一个set

type Set struct {
	m   map[int]struct{} // 为什么我们要使用空结构体，因为空结构体不占用内存
	len int
	sync.RWMutex
}

// NewSet 新建一个set
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}

// Add 增加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()

	s.m[item] = struct{}{}
	s.len = len(s.m)
}

//Remove 移除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()

	if s.len == 0 {
		return
	}
	// 从字典中删除
	delete(s.m, item)
	// 计算长度
	s.len = len(s.m)
}

// Has 判断一个元素是否在set中
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.m[item]
	return ok
}

// Len 获取set的长度
func (s *Set) Len() int {
	return s.len
}

//IsEmpty 判断set是否为空
func (s *Set) IsEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}

// Clear 清空set
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()

	s.m = make(map[int]struct{})
	s.len = 0
}

// List 将 Set 转化为 Slice
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()

	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}