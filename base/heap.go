package base

import "github.com/skeletongo/datastructure/heap"

// HeapSort 堆排序
// 将原数组原地排序
func HeapSort(arr []interface{}, compare func(a, b interface{}) int) {
	n := len(arr)
	if n < 2 {
		return
	}

	i := (n - 2) / 2
	for ; i >= 0; i-- {
		siftDown(arr, i, compare)
	}
	arr[0], arr[n-1] = arr[n-1], arr[0]

	for i := n - 2; i > 0; i-- {
		siftDown(arr[:i+1], 0, compare)
		arr[0], arr[i] = arr[i], arr[0]
	}
}

func siftDown(arr []interface{}, i int, compare func(a, b interface{}) int) {
	n := len(arr)
	e := arr[i]
	for {
		j := 2*i + 1
		if j >= n || j < 0 { // j < 0,当 2*i+1 整数溢出时
			break
		}
		if ri := j + 1; ri < n && compare(arr[j], arr[ri]) < 0 {
			j = ri
		}
		if compare(e, arr[j]) >= 0 {
			break
		}
		arr[i] = arr[j]
		i = j
	}
	arr[i] = e
}

// IndexHeapSort 堆排序
// 元素入堆再取出的方式
func HeapSort2(arr []interface{}, compare func(a, b interface{}) int) {
	h := heap.New(compare)
	for _, v := range arr {
		h.Add(v)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = h.ExtractMax()
	}
}
