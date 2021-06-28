package base

/*冒泡排序*/

func BubbleSort(arr []interface{}, compare func(a, b interface{}) int) {
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 0; j < len(arr)-1-i; j++ {
			if compare(arr[j], arr[j+1]) > 0 {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}

		if !flag { // 已排好序
			break
		}
	}
}

func BubbleSort2(arr []interface{}, compare func(a, b interface{}) int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if compare(arr[j], arr[j-1]) < 0 {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func BubbleSort3(arr []interface{}, compare func(a, b interface{}) int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if compare(arr[j], arr[j+1]) > 0 {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
