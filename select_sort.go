package algorithm

/*
 选择排序 时间复杂度O(n^2)
 排序分成两层循环，第一层循环从第一个元素(用i表示下标)开始向后遍历，第二层循环从外层循环的下一个元素(用j表示下标)开始向后遍历
 在第二层循环开始前声明一个变量(用m表示)，记录从第二层循环中搜索到的最小元素的下标，第二层循环体中搜索arr[j]<arr[m]的情况并将j赋值给m
 经过第二层循环之后将最小元素(m对应的元素)和第一层循环的当前元素(i对应的元素)交换
*/
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		m := i // 记录最小元素所在下标
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[m] {
				m = j
			}
		}
		arr[i], arr[m] = arr[m], arr[i]
	}
}
