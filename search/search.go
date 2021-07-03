package search

// Search 二分查找
// v 查找的值
// 返回 v 在数组 arr 中的索引，没有找到返回-1
func Search(arr []float64, v float64) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
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

// Search2 二分查找(递归)
// v 查找的值
// 返回 v 在数组 arr 中的索引，没有找到返回-1
func Search2(arr []float64, v float64) int {
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

// Floor 在arr中搜索小于等于v的最大值所在的索引，没有找到返回-1
func Floor(arr []float64, v float64) int {
	l := 0
	r := len(arr) - 1

	ret := -1
	for l <= r {
		mid := l + (r-l)/2
		if v < arr[mid] {
			r = mid - 1
		} else if v > arr[mid] {
			ret = mid
			l = mid + 1
		} else {
			return mid
		}
	}
	return ret
}

// Ceil 在arr中搜索大于等于v的最小值所在的索引，没有找到返回-1
func Ceil(arr []float64, v float64) int {
	l := 0
	r := len(arr) - 1

	ret := -1
	for l <= r {
		mid := l + (r-l)/2
		if v < arr[mid] {
			ret = mid
			r = mid - 1
		} else if v > arr[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return ret
}
