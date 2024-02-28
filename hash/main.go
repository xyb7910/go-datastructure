package hash

import (
	"fmt"
	"github.com/cespare/xxhash"
	"math"
	"sync"
)

const (
	// 扩容因子
	expandFactor = 0.75
)

// 键值对
type keyPairs struct {
	key   string
	value interface{}
	next  *keyPairs
}

// HashMap 哈希表
type HashMap struct {
	array        []*keyPairs
	len          int
	capacity     int
	capacityMask int
	lock         sync.Mutex
}

// NewHashMap 初始化哈希表
func NewHashMap(capacity int) *HashMap {
	// 默认容积为2的幂
	defaultCapacity := 1 << 4
	if capacity <= defaultCapacity {
		capacity = defaultCapacity
	} else {
		capacity = 1 << int(math.Ceil(math.Log2(float64(capacity))))
	}

	// 新建一个哈希表
	hashtable := new(HashMap)
	hashtable.capacity = capacity
	hashtable.capacityMask = capacity - 1
	return hashtable
}

// Len 返回哈希表中键值对的个数
func (hashtable *HashMap) Len() int {
	return hashtable.len
}

// value 计算哈希值
var value = func(key []byte) uint64 {
	h := xxhash.New()
	h.Write(key)
	return h.Sum64()
}

// hashIndex 计算哈希值并获取下标
func (hashtable *HashMap) hashIndex(key string, mask int) int {
	// 计算哈希值
	hash := value([]byte(key))
	index := hash & uint64(mask)
	return int(index)
}

// Put 插入键值对
func (hashtable *HashMap) Put(key string, value interface{}) {
	hashtable.lock.Lock()
	defer hashtable.lock.Unlock()

	// 获取下标
	index := hashtable.hashIndex(key, hashtable.capacityMask)
	// 此下标在哈希表中的值
	element := hashtable.array[index]
	if element == nil {
		// 此下标没有元素，则插入
		hashtable.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		// 此下标已经有元素，则插入到上一个元素的后面
		var lastPairs *keyPairs

		for element != nil {
			if element.key == key {
				element.value = value
				return
			}
			lastPairs = element
			element = element.next
		}

		// 找不到元素，则插入到最后
		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}
	}
	// 长度加一
	newLen := hashtable.len + 1

	// 计算扩容因子，如果长度大于容积的75%，则扩容
	if float64(newLen)/float64(hashtable.capacity) >= expandFactor {
		// 新建一个原来两倍大小的哈希表
		newhashtable := new(HashMap)
		newhashtable.array = make([]*keyPairs, hashtable.capacity*2)
		newhashtable.capacity = hashtable.capacity * 2
		newhashtable.capacityMask = newhashtable.capacity*2 - 1

		// 遍历原哈希表，将元素插入到新哈希表
		for _, pairs := range hashtable.array {
			for pairs != nil {
				newhashtable.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		hashtable.array = newhashtable.array
		hashtable.capacity = newhashtable.capacity
		hashtable.capacityMask = newhashtable.capacityMask
	}
	hashtable.len = newLen
}

// Get 获取键值对
func (hashtable *HashMap) Get(key string) (value interface{}, ok bool) {
	hashtable.lock.Lock()
	defer hashtable.lock.Unlock()

	// 获取下标
	index := hashtable.hashIndex(key, hashtable.capacityMask)

	// 此下标在哈希表中的值
	element := hashtable.array[index]

	// 遍历元素，如果元素的key等于key，则返回
	for element != nil {
		if element.key == key {
			return element.value, true
		}
		element = element.next
	}
	return nil, false
}

// Delete 删除键值对
func (hashtable *HashMap) Delete(key string) {
	hashtable.lock.Lock()
	defer hashtable.lock.Unlock()

	// 获取下标
	index := hashtable.hashIndex(key, hashtable.capacityMask)

	// 此下标在哈希表中的值
	element := hashtable.array[index]

	// 如果为空链表，则直接返回
	if element == nil {
		return
	}

	// 如果第一个元素的key等于key，则删除
	if element.key == key {
		hashtable.array[index] = element.next
		hashtable.len--
		return
	}

	// 下一个键值对
	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			element.next = nextElement.next
			hashtable.len--
			return
		}
		element = nextElement
		nextElement = nextElement.next
	}
}

// Range 遍历哈希表
func (hashtable *HashMap) Range() {
	hashtable.lock.Lock()
	defer hashtable.lock.Unlock()

	for _, pairs := range hashtable.array {
		for pairs != nil {
			fmt.Println(pairs.key, pairs.value)
			pairs = pairs.next
		}
	}
	fmt.Println("len:", hashtable.len)
}
