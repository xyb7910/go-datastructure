package main

type Ring struct {
	next, prev *Ring
	Value      interface{}
}

// 初始化循环链表
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// New 创建一个链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

/*
获取下一个节点
获取上一个节点
*/

func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

// 当 n 为负数的时候，表示向前移动，当 n 为正数的时候，表示向后移动

func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 往节点A， 连接一个节点，并且返回之前节点A的后驱节点

func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 删除节点后的第 n 个节点

func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// 获取链表长度

func (r *Ring) Len() int {
	n := 0
	if r != nil {
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}
