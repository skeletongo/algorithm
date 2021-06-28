package base

// SelectSort 选择排序
// compare 元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func SelectSort(arr []interface{}, compare func(a, b interface{}) int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if compare(arr[j], arr[min]) < 0 {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
