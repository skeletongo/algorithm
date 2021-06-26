package base

/*
 希尔排序
 希尔排序基于插入排序，时间复杂度和希尔增量有关，排序时间优于直接插入排序
*/

// ShellSort 希尔排序
// 递减增量 h = l/2 元素个数除以2之后每次缩小一半
func ShellSort(arr []int, l, r int) {
	for h := (r - l + 1) / 2; h > 0; h /= 2 {
		for i := l + h; i <= r; i++ {
			e := arr[i]
			j := i
			for ; j-h >= l && arr[j-h] > e; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
	}
}

// ShellSort2 希尔排序
// 递减增量 h = 3*d +1 	1=3*0+1 4=3*1+1 13=3*4+1 40=3*13+1 121=3*40+1 364=3*121+1
func ShellSort2(arr []int, l, r int) {
	var h int
	n := r - l + 1
	for d := 0; d < n; d = 3*d + 1 {
		h = d
	}
	for ; h > 0; h = (h - 1) / 3 {
		for i := l + h; i <= r; i++ {
			e := arr[i]
			j := i
			for ; j-h >= l && arr[j-h] > e; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
	}
}
