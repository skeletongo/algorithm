package base

// 归并排序自顶向下
func MergeSort(arr []int, n int) {
	mergeSort(arr, 0, n-1)
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	m := l + (r-l)/2 // 防止整数溢出
	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)
	merge(arr, l, m, r)
}

func merge(arr []int, l, m, r int) {
	data := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		data[i-l] = arr[i]
	}

	i := l
	j := m + 1

	for k := l; k <= r; k++ {
		if j > r {
			arr[k] = data[i-l]
			i++
		} else if i > m {
			arr[k] = data[j-l]
			j++
		} else if data[i-l] < data[j-l] {
			arr[k] = data[i-l]
			i++
		} else {
			arr[k] = data[j-l]
			j++
		}
	}
}

// 归并排序自底向上
func MergeSortBU(arr []int, n int) {
	for sz := 1; sz < n; sz += sz { // 每组中的元素个数
		for i := 0; i+sz < n; i += 2 * sz { // 待合并的两组数据
			merge(arr, i, i+sz-1, min(i+2*sz-1, n-1))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
