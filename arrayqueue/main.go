package arrayqueue

import "sync"

type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

// Add 添加一个元素給队列
func (queue *ArrayQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.array = append(queue.array, v)

	queue.size++
}

// Remove 移除队列的第一个元素
func (queue *ArrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("Queue is empty")
	}

	v := queue.array[0]

	/* 原地移动，但缩容后空间不会被释放
	for i := 1; i < queue.size; i++ {
		queue.array[i-1] = queue.array[i]
	}
	// 缩容后，数组长度减1
	queue.array = queue.array[:queue.size-1]
	*/
	newArray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray
	queue.size--
	return v
}
