package base

import "container/list"

// InsertSort 插入排序
// compare 元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func InsertSort(arr []interface{}, compare func(a, b interface{}) int) {
	for i := 1; i < len(arr); i++ {
		e := arr[i]
		j := i
		for ; j > 0 && compare(arr[j-1], e) > 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

// InsertSortList 链表插入排序
// l 需要排序的链表
// s 规定链表的第一个元素的索引为0，第二个元素索引为1，元素索引依次递增，s表示从索引为s的元素开始排序，包含s位置
// n 表示从索引s开始一共要排序的元素总数
// compare 元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func InsertSortList(l *list.List, s, n int, compare func(a, b interface{}) int) {
	e := l.Front()
	for i := 0; i < s && e != nil; i++ {
		e = e.Next()
	}
	InsertSortListByElement(e, n, compare)
}

// InsertSortListByElement 链表插入排序
// n 从当前元素开始需要排序的元素总数
// compare 元素大小比较方法
// 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func InsertSortListByElement(e *list.Element, n int, compare func(a, b interface{}) int) {
	if e == nil {
		return
	}

	front := e.Prev()
	e = e.Next()
	for i := 0; e != nil && i < n-1; i++ {
		v := e.Value
		ej := e
		for ; ej.Prev() != front && compare(ej.Prev().Value, v) > 0; ej = ej.Prev() {
			ej.Value = ej.Prev().Value
		}
		ej.Value = v
		e = e.Next()
	}
}
