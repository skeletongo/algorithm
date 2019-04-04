package algorithm

/*
 冒泡排序 最差时间复杂度O(n^2) 最佳时间复杂度O(n)
 外层循环控制冒泡次数(n个元素外层循环次数为n-1次)
 内层循环做元素比较和交换操作，最多交换“元素个数减去外层已循环次数再减1”次
*/
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func BubbleSort2(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func BubbleSort3(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		flag := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
