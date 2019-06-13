package base

// 选择排序
func InsertSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		e := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

/*
 希尔排序
 希尔排序基于插入排序，时间复杂度和希尔增量有关，排序时间优于直接插入排序
*/
// 递减增量 h = l/2 元素个数除以2之后每次缩小一半
func ShellSort(arr []int) {
	for h := len(arr) / 2; h > 0; h /= 2 {
		for i := h; i < len(arr); i++ {
			e := arr[i]
			j := i
			for ; j-h >= 0 && arr[j-h] > e; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
	}
}

// 递减增量 h = 3*d +1 	1=3*0+1 4=3*1+1 13=3*4+1 40=3*13+1 121=3*40+1 364=3*121+1
func ShellSort2(arr []int) {
	var h int
	for d := 0; d < len(arr); d = 3*d + 1 {
		h = d
	}
	for ; h > 0; h = (h - 1) / 3 {
		for i := h; i < len(arr); i++ {
			e := arr[i]
			j := i
			for j = i; j-h >= 0 && arr[j-h] > e; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
	}
}
