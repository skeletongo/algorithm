package base

import (
	"container/list"
)

// MergeSort 归并排序
// 自顶向下,递归方式
// compare 元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func MergeSort(arr []interface{}, compare func(a, b interface{}) int) {
	mergeSort(arr, 0, len(arr)-1, compare)
}

func mergeSort(arr []interface{}, l, r int, compare func(a, b interface{}) int) {
	if r-l <= 15 {
		InsertSort(arr[l:r+1], compare)
		return
	}

	m := l + (r-l)/2 // 防止整数溢出
	mergeSort(arr, l, m, compare)
	mergeSort(arr, m+1, r, compare)
	if compare(arr[m], arr[m+1]) > 0 {
		merge(arr, l, m, r, compare)
	}
}

func merge(arr []interface{}, l, m, r int, compare func(a, b interface{}) int) {
	data := make([]interface{}, r-l+1)
	for i := l; i <= r; i++ {
		data[i-l] = arr[i]
	}

	i := l
	j := m + 1

	for k := l; k <= r; k++ {
		if j > r {
			arr[k] = data[i-l]
			i++
		} else if i > m {
			arr[k] = data[j-l]
			j++
		} else if compare(data[i-l], data[j-l]) <= 0 {
			arr[k] = data[i-l]
			i++
		} else {
			arr[k] = data[j-l]
			j++
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MergeSortBU 归并排序
// 自底向上,非递归方式
// compare 元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func MergeSortBU(arr []interface{}, compare func(a, b interface{}) int) {
	for sz := 1; sz < len(arr); sz += sz { // 每组中的元素个数
		for i := 0; i+sz < len(arr); i += 2 * sz { // 待合并的两组数据
			rr := min(i+2*sz-1, len(arr)-1)
			if rr-i <= 15 {
				InsertSort(arr[i:rr+1], compare)
			} else if compare(arr[i+sz-1], arr[i+sz]) > 0 {
				merge(arr, i, i+sz-1, rr, compare)
			}
		}
	}
}

// MergeSortList 对链表进行归并排序
// l 需要排序的链表
// s 规定链表的第一个元素的索引为0，第二个元素索引为1，元素索引依次递增，s表示从索引为s的元素开始排序，包含s位置
// n 表示从索引s开始一共要排序的元素总数
// compare 元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func MergeSortList(l *list.List, s, n int, compare func(a, b interface{}) int) {
	e := l.Front()
	for i := 0; i < s && e != nil; i++ {
		e = e.Next()
	}
	MergeSortListByElement(e, n, compare)
}

// MergeSortListByElement 对链表进行归并排序
// n 从当前元素开始需要排序的元素总数
// compare 元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func MergeSortListByElement(e *list.Element, n int, compare func(a, b interface{}) int) {
	if e == nil {
		return
	}
	front := e
	for sz := 1; sz < n; sz += sz { // 每组中的元素个数
		for e = front; e != nil; {
			ei := e
			var data []interface{}
			for i := 0; i < 2*sz && e != nil; i++ {
				data = append(data, e.Value)
				e = e.Next()
			}

			if len(data) <= sz {
				break
			}

			if compare(data[sz-1], data[sz]) <= 0 {
				continue
			}

			if len(data) <= 15 {
				InsertSortListByElement(ei, len(data), compare)
				continue
			}

			i := 0
			j := sz
			r := len(data) - 1

			for k := 0; k <= r; k++ {
				if j > r {
					ei.Value = data[i]
					i++
				} else if i > sz-1 {
					ei.Value = data[j]
					j++
				} else if compare(data[i], data[j]) <= 0 {
					ei.Value = data[i]
					i++
				} else {
					ei.Value = data[j]
					j++
				}
				ei = ei.Next()
			}
		}
	}
}
