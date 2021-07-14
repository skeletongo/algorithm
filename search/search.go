package search

// Search 二分查找
// v 查找的值
// 返回 v 在数组 arr 中的任意一个索引，没有找到返回-1
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

// Floor 在数组arr中搜索值等于v的元素的索引，如果有多个，返回第一个
// 如果没有则返回小于v的最大值元素的索引，如果有多个，返回最大索引，如果还没有就返回-1
func Floor(arr []float64, v float64) int {
	l := -1
	r := len(arr) - 1

	for l < r {
		//
		mid := l + (r-l+1)/2
		if arr[mid] >= v {
			r = mid - 1
		} else {
			l = mid
		}
	}

	// 如果该索引+1就是搜索值本身, 该索引+1即为返回值
	if l+1 < len(arr) && arr[l+1] == v {
		return l + 1
	}

	// 否则, 该索引即为返回值
	return l
}

// Ceil 在数组arr中搜索值等于v的元素的索引，如果有多个，返回最后一个
// 如果没有则返回大于v的最小值元素的索引，如果有多个，返回最小索引，如果还没有就返回-1
func Ceil(arr []float64, v float64) int {
	l := 0
	r := len(arr)

	for l < r {
		//
		mid := l + (r-l)/2
		if arr[mid] <= v {
			l = mid + 1
		} else {
			r = mid
		}
	}

	// 如果该索引-1就是搜索值本身, 该索引-1即为返回值
	if r-1 >= 0 && arr[r-1] == v {
		return r - 1
	}

	if len(arr) == r {
		return -1
	}
	return r
}
