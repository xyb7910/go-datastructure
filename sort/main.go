package sort

/*
这个排序的原理大家可以理解为小鱼吐泡泡，每次大泡泡会在最上边。

在实际的生活中根本不用，效率太低了。了解思路即可
*/

// BubbleSort 冒泡排序
func BubbleSort(list []int) {
	n := len(list)
	didSwap := true
	// 第一次遍历为控制排序的轮数 N - 1
	for i := n - 1; i > 0; i-- {
		// 每一轮遍历完之后都会有一个元素的位置确定，
		/*
			为什么要比较到 i 位呢，因为每一次遍历后，第 i 位的元素是已经确定顺序的
		*/
		for j := 0; j < i; j++ {
			// 如果前边的元素大于后边的元素，则进行交换
			if list[j] > list[j+1] {
				didSwap = true
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
		// 如果没有进行交换，则说明已经排序完成
		if !didSwap {
			return
		}
	}
}

/*
原理和我们平时玩扑克拍时候，对牌进行整理一个原理哈
在每次打扑克的时候，会习惯的从左到右扫描，然后将最小的牌放在最左边，
然后在从最左边的第二张牌开始继续，从左到右开始扫描第二小牌，放在最小牌的右边，
如此，重复下去，直到所有牌都整理完成
*/

// SelectSort 选择排序
func SelectSort(list []int) {
	n := len(list)
	// 总共进行 n - 1 轮排序
	for i := 0; i < n-1; i++ {
		// 每次从第 i 个元素开始，选取最小的元素
		min := list[i] // 选取的最小元素
		minIndex := i  // 最小元素的位置
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		// 如果最小元素的位置不等于 i，则进行交换
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

/*
对于选择排序的优化，普通的选择排序，每次我们之后将最小的元素找出来，
这样的做法，需要我们进行 N - 1 次

优化的思路为：
	每次除了寻找最小元素之外，我们还可以找出最大的元素，
	然后分别于前面和后面的元素进行交换，这样，循环次数直接减少一半。

虽然我进行了优化，但实际证明，这个算法的效率还是不高，所以在工程中，不要用！！！
*/

// SelectProSort 选择排序的优化
func SelectProSort(list []int) {
	n := len(list)
	// 总共会循环 n / 2 次
	for i := 0; i < n/2; i++ {
		minIndex := i
		maxIndex := i
		// 每一次循环，我们会找到最小的元素，然后找到最大的元素
		for j := i + 1; j < n-i; j++ {
			// 寻找最大的元素下标
			if list[j] > list[maxIndex] {
				maxIndex = j // 更新最大的元素下标
				continue
			}
			// 寻找最小的元素下标
			if list[j] < list[minIndex] {
				minIndex = j // 更新最小的元素下标
			}
		}
		// 如果最大元素是开头的元素，但最小元素不是最尾部的元素
		// 则先将最大元素与最尾部元素进行交换
		if maxIndex == i && minIndex == i {
			list[i], list[maxIndex] = list[maxIndex], list[n-i-1]
			// 然后将最小元素与开头的元素进行交换
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == n-i-1 {
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			// 否则，先将最小元素放在开头，然后将最大元素放到最尾部
			list[i], list[minIndex] = list[minIndex], list[i]
			list[maxIndex], list[n-i-1] = list[n-i-1], list[maxIndex]
		}
	}
}

/*
我们也用玩扑克举例子吧。
有些人在玩扑克的时候，习惯从第二张牌开始，和第一张牌进行比较
第二张牌如果比第一张牌小，则将第二张牌放到第一张牌的左边，，这样两个牌顺序就对了
接着从第三张牌开始，将它插入到已经牌好序的前二张牌里，形成三张有序的牌
以此类推...
*/

// InsertSort 插入排序
func InsertSort(list []int) {
	n := len(list)
	// 总共进行 n - 1 轮排序,并且记录了下标
	for i := 1; i <= n-1; i++ {
		dealNum := list[i] // 要处理的元素
		j := i - 1         // 要处理的元素的位置
		// 如果发现左边的元素比当前元素大，则进行处理
		if dealNum < list[j] {
			// 一直往左找，比待排序的元素大的元素，我们直接将其向后移动
			for ; j >= 0 && dealNum < list[j]; j-- {
				list[j+1] = list[j]
			}
			// 将带排序元素放到正确位置
			list[j+1] = dealNum
		}
	}
}
