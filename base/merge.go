package base

func mergeSort(arr []int, l, r int) {
	if r-l < 15 {
		InsertSort(arr, l, r)
		return
	}

	m := l + (r-l)/2 // 防止整数溢出
	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)
	if arr[m] > arr[m+1] {
		merge(arr, l, m, r)
	}
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

// MergeSort 归并排序
// 自顶向下,递归方式
func MergeSort(arr []int, l, r int) {
	mergeSort(arr, l, r)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MergeSortBU 归并排序
// 自底向上,非递归方式
func MergeSortBU(arr []int, l, r int) {
	for sz := l + 1; sz <= r; sz += sz { // 每组中的元素个数
		for i := l; i+sz <= r; i += 2 * sz { // 待合并的两组数据
			rr := min(i+2*sz-1, r)
			if rr-i < 15 {
				InsertSort(arr, i, rr)
			} else if arr[i+sz-1] > arr[i+sz] {
				merge(arr, i, i+sz-1, rr)
			}
		}
	}
}
