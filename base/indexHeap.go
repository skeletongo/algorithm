package base

// 索引堆
type IndexMaxHeap struct {
	// i 为堆索引
	// reverse[indexes[i]] = i
	data    []int
	indexes []int
	reverse []int

	arr []int // 可用data位置
}

func (i *IndexMaxHeap) GetSize() int {
	return len(i.indexes)
}

func (i *IndexMaxHeap) IsEmpty() bool {
	return len(i.indexes) == 0
}

func (i *IndexMaxHeap) contains(index int) bool {
	if index < 0 || index >= len(i.data) {
		panic("param error")
	}
	return i.reverse[index] != -1
}

// 根据元素索引查询元素的值
func (i *IndexMaxHeap) Get(index int) int {
	if !i.contains(index) {
		panic("Invalid index")
	}
	return i.data[index]
}

func (i *IndexMaxHeap) Change(index int, n int) {
	if !i.contains(index) {
		panic("Invalid index")
	}

	i.data[index] = n
	j := i.reverse[index]
	i.shiftUp(j)
	i.shiftDown(j)
}

// 往堆中插入一个元素，返回这个元素的查询索引
func (i *IndexMaxHeap) Insert(n int) int {
	// 获取可用位置
	index := -1
	if len(i.arr) > 0 {
		index = i.arr[len(i.arr)-1]
		i.arr = i.arr[:len(i.arr)-1]
	}

	if index == -1 { // 追加元素
		i.data = append(i.data, n)
		index = len(i.data) - 1
		i.reverse = append(i.reverse, len(i.indexes))
	} else { // 覆盖已取出元素的值
		i.data[index] = n
		i.reverse[index] = len(i.indexes)
	}

	i.indexes = append(i.indexes, index)
	i.shiftUp(len(i.indexes) - 1)
	return index
}

func (i *IndexMaxHeap) ExtractMax() int {
	if i.GetSize() == 0 {
		panic("no data")
	}
	e := i.data[i.indexes[0]]
	i.arr = append(i.arr, i.indexes[0]) // 记录已取出元素的data下标

	i.indexes[0], i.indexes[len(i.indexes)-1] = i.indexes[len(i.indexes)-1], i.indexes[0]
	i.reverse[i.indexes[0]] = 0
	i.reverse[i.indexes[len(i.indexes)-1]] = -1
	i.indexes = i.indexes[:len(i.indexes)-1]
	i.shiftDown(0)
	return e
}

// 时间复杂度O(n)
func (i *IndexMaxHeap) Heapify(arr []int) {
	for k, v := range arr {
		i.data = append(i.data, v)
		i.indexes = append(i.indexes, k)
		i.reverse = append(i.reverse, k)
	}
	for k := i.parent(len(arr) - 1); k >= 0; k-- {
		i.shiftDown(k)
	}
}

func (i *IndexMaxHeap) shiftUp(k int) {
	for ; k > 0 && i.data[i.indexes[i.parent(k)]] < i.data[i.indexes[k]]; k = i.parent(k) {
		i.indexes[i.parent(k)], i.indexes[k] = i.indexes[k], i.indexes[i.parent(k)]
		i.reverse[i.indexes[i.parent(k)]] = i.parent(k)
		i.reverse[i.indexes[k]] = k
	}
}

func (i *IndexMaxHeap) shiftDown(k int) {
	size := i.GetSize()
	for i.leftChild(k) < size {
		j := i.leftChild(k)
		if j+1 < size && i.data[i.indexes[j]] < i.data[i.indexes[j+1]] {
			j++
		}

		if i.data[i.indexes[k]] >= i.data[i.indexes[j]] {
			break
		}

		i.indexes[k], i.indexes[j] = i.indexes[j], i.indexes[k]
		i.reverse[i.indexes[k]] = k
		i.reverse[i.indexes[j]] = j
		k = j
	}
}

func (i *IndexMaxHeap) parent(k int) int {
	return (k - 1) / 2
}

func (i *IndexMaxHeap) leftChild(k int) int {
	return 2*k + 1
}

func (i *IndexMaxHeap) rightChild(k int) int {
	return 2*k + 2
}

// 堆排序1
func IndexHeapSort1(arr []int) {
	maxHeap := IndexMaxHeap{}
	for _, v := range arr {
		maxHeap.Insert(v)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
}

// 堆排序2
func IndexHeapSort2(arr []int) {
	maxHeap := IndexMaxHeap{}
	maxHeap.Heapify(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
}