package base

import (
	"math/rand"
	"time"
)

var R = rand.New(rand.NewSource(time.Now().UnixNano()))

// 快速排序
func QuickSort(arr []int, n int) {
	quickSort2(arr, 0, n-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partition2(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

// 基础版
func partition1(arr []int, l, r int) int {
	e := arr[l]

	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < e {
			arr[i], arr[j+1] = arr[j+1], arr[i]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// 随机数优化
func partition2(arr []int, l, r int) int {
	t := l + R.Intn(r-l+1)
	arr[l], arr[t] = arr[t], arr[l]

	return partition3(arr, l, r)
}

// 双路快速排序
func partition3(arr []int, l, r int) int {
	e := arr[l]

	i := l + 1
	j := r
	for {
		for i <= r && arr[i] < e {
			i++
		}
		for j > l && arr[j] > e {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// 三路快速排序
func quickSort2(arr []int, l, r int) {
	if l >= r {
		return
	}

	// partition
	t := l + R.Intn(r-l+1)
	arr[l], arr[t] = arr[t], arr[l]

	e := arr[l]

	lt := l
	i := l + 1
	gt := r + 1

	for i != gt {
		if arr[i] < e {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			lt++
			i++
		} else if arr[i] > e {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]

	quickSort2(arr, l, lt-1)
	quickSort2(arr, gt, r)
}
