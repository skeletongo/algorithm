package base

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// QuickSort 三路快速排序
// compare 元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func QuickSort(arr []interface{}, compare func(a, b interface{}) int) {
	quickSort(arr, compare)
}

func quickSort(arr []interface{}, compare func(a, b interface{}) int) {
	if len(arr) <= 15 {
		InsertSort(arr, compare)
		return
	}

	// 随机取一个值作为中间值
	t := rand.Intn(len(arr))
	arr[0], arr[t] = arr[t], arr[0]

	e := arr[0]

	lt := 0
	i := 1
	gt := len(arr)

	for i < gt {
		if compare(arr[i], e) < 0 {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			lt++
			i++
		} else if compare(arr[i], e) > 0 {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[0], arr[lt] = arr[lt], arr[0]

	quickSort(arr[:lt], compare)
	quickSort(arr[gt:], compare)
}
