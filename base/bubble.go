package base

/*冒泡排序*/

func BubbleSort(arr []int, l, r int) {
	for i := l; i < r; i++ {
		flag := false
		for j := l; j < r-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}

		if !flag { // 已排好序
			break
		}
	}
}

func BubbleSort2(arr []int, l, r int) {
	for i := l; i < r; i++ {
		for j := r; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func BubbleSort3(arr []int, l, r int) {
	for i := l; i < r; i++ {
		for j := l; j < r-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
