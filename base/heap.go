package base

// 最大堆
type MaxHeap struct {
	arr []int
}

func (m *MaxHeap) GetSize() int {
	return len(m.arr)
}

func (m *MaxHeap) IsEmpty() bool {
	return len(m.arr) == 0
}

func (m *MaxHeap) Insert(n int) {
	m.arr = append(m.arr, n)
	m.shiftUp(m.GetSize() - 1)
}

func (m *MaxHeap) ExtractMax() int {
	if m.GetSize() == 0 {
		panic("no data")
	}
	e := m.arr[0]
	m.arr[0] = m.arr[len(m.arr)-1]
	m.arr = m.arr[:len(m.arr)-1]
	m.shiftDown(0)
	return e
}

// 时间复杂度O(n)
func (m *MaxHeap) Heapify(arr []int) {
	for _, v := range arr {
		m.arr = append(m.arr, v)
	}
	for k := m.parent(len(arr) - 1); k >= 0; k-- {
		m.shiftDown(k)
	}
}

func (m *MaxHeap) shiftUp(k int) {
	for ; k > 0 && m.arr[m.parent(k)] < m.arr[k]; k = m.parent(k) {
		m.arr[m.parent(k)], m.arr[k] = m.arr[k], m.arr[m.parent(k)]
	}
}

func (m *MaxHeap) shiftDown(k int) {
	size := m.GetSize()
	for m.leftChild(k) < size {
		j := m.leftChild(k)
		if j+1 < size && m.arr[j] < m.arr[j+1] {
			j++
		}

		if m.arr[k] >= m.arr[j] {
			break
		}

		m.arr[k], m.arr[j] = m.arr[j], m.arr[k]
		k = j
	}
}

func (m *MaxHeap) parent(k int) int {
	return (k - 1) / 2
}

func (m *MaxHeap) leftChild(k int) int {
	return 2*k + 1
}

func (m *MaxHeap) rightChild(k int) int {
	return 2*k + 2
}

// 堆排序1
func HeapSort1(arr []int) {
	maxHeap := MaxHeap{}
	for _, v := range arr {
		maxHeap.Insert(v)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
}

// 堆排序2
func HeapSort2(arr []int) {
	maxHeap := MaxHeap{}
	maxHeap.Heapify(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
}

// 原地堆排序（不需要开辟新空间）
func HeapSort(arr []int) {
	for i := (len(arr) - 1) / 2; i >= 0; i-- {
		shiftDown(arr, i)
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown(arr[:i], 0)
	}
}

func shiftDown(arr []int, k int) {
	size := len(arr)
	for 2*k+1 < size {
		j := 2*k + 1
		if j+1 < size && arr[j] < arr[j+1] {
			j++
		}

		if arr[k] >= arr[j] {
			break
		}

		arr[k], arr[j] = arr[j], arr[k]
		k = j
	}
}
