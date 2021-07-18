package search

import "sort"

// 二分搜索标准库sort包已实现就是Search方法，但是没有Floor，Ceil方法，不过用Search方法就可以实现
//
// 包中的所有方法中的compare参数都是元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b

// SearchFunc 在升序的切片arr中搜索特定值x，返回x值所在的索引，如果有多个返回最小索引，如果没有搜索到返回-1
func SearchFunc(arr []interface{}, compare func(a, b interface{}) int, x interface{}) int {
	i := sort.Search(len(arr), func(i int) bool {
		return compare(arr[i], x) >= 0
	})
	if i < len(arr) && compare(arr[i], x) == 0 {
		return i
	}
	return -1
}

// FloorFunc 在升序的切片arr中搜索比特定值x小的所有值中最大值的索引，如果有多个返回最大索引，如果没有搜索到返回-1
func FloorFunc(arr []interface{}, compare func(a, b interface{}) int, x interface{}) int {
	i := sort.Search(len(arr), func(i int) bool {
		return compare(arr[i], x) >= 0
	})
	if i-1 >= 0 {
		return i - 1
	}
	return -1
}

// CeilFunc 在升序的切片arr中搜索比特定值x大的所有值中最小值的索引，如果有多个返回最小索引，如果没有搜索到返回-1
func CeilFunc(arr []interface{}, compare func(a, b interface{}) int, x interface{}) int {
	i := sort.Search(len(arr), func(i int) bool {
		return compare(arr[i], x) > 0
	})
	if i < len(arr) {
		return i
	}
	return -1
}

// Search 二分查找
// v 查找的值
// 返回 v 在数组 arr 中的任意一个索引，没有找到返回-1
//Deprecated
func Search(arr []float64, v float64) int {
	l := 0
	r := len(arr) - 1

	// 在arr[l...r]之中查找目标值
	for l <= r {
		// mid = (l + r)/2 防止极端情况下的整形溢出，使用下面的形式
		mid := l + (r-l)/2
		if v < arr[mid] {
			r = mid - 1
		} else if v > arr[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// SearchR 二分查找(递归)
// v 查找的值
// 返回 v 在数组 arr 中的任意一个索引，没有找到返回-1
//Deprecated
func SearchR(arr []float64, v float64) int {
	l := 0
	r := len(arr) - 1
	return searchR(arr, l, r, v)
}

func searchR(arr []float64, l, r int, v float64) int {
	if r < l {
		return -1
	}
	mid := l + (r-l)/2
	if v < arr[mid] {
		return searchR(arr, l, mid-1, v)
	}
	if v > arr[mid] {
		return searchR(arr, mid+1, r, v)
	}
	return mid
}
