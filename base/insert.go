package base

// InsertSort 插入排序
func InsertSort(arr []int, l, r int) {
	for i := l + 1; i <= r; i++ {
		e := arr[i]
		j := i
		for ; j > l && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
