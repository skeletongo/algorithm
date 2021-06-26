package base

// SelectSort 选择排序
func SelectSort(arr []int, l, r int) {
	for i := l; i <= r; i++ {
		min := i
		for j := i + 1; j <= r; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
