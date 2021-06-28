package base

import (
	"container/list"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func compare(a, b interface{}) int {
	return a.(int) - b.(int)
}

func newInterfaceSlice(l, r int, n int) []interface{} {
	ret := make([]interface{}, n)
	for i := 0; i < n; i++ {
		ret[i] = l + rand.Intn(r-l+1)
	}
	return ret
}

func copyInterfaceSlice(data []int) []interface{} {
	ret := make([]interface{}, len(data))
	for i := 0; i < len(data); i++ {
		ret[i] = data[i]
	}
	return ret
}

func newIntSlice(l, r int, n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = l + rand.Intn(r-l+1)
	}
	return ret
}

func newList(l, r int, n int) *list.List {
	ls := list.New()
	for i := 0; i < n; i++ {
		ls.PushBack(l + rand.Intn(r-l+1))
	}
	return ls
}

func sortTest(data []interface{},
	f func([]interface{}, func(interface{}, interface{}) int),
	compare func(a, b interface{}) int) bool {
	f(data, compare)
	for i := 1; i < len(data); i++ {
		if compare(data[i-1], data[i]) > 0 {
			return false
		}
	}
	return true
}

func sortListTest(l *list.List,
	f func(*list.List, int, int, func(a, b interface{}) int),
	compare func(a, b interface{}) int) bool {
	flag := false
	f(l, 0, l.Len(), compare)
	for e := l.Front().Next(); e != nil; e = e.Next() {
		//fmt.Print(e.Value, ",")
		if compare(e.Prev().Value, e.Value) > 0 {
			flag = true
		}
	}
	return !flag
}

func sortTimeTest(name string, data []interface{},
	f func([]interface{}, func(interface{}, interface{}) int),
	compare func(a, b interface{}) int) {
	t := time.Now()
	f(data, compare)
	fmt.Printf("%10s: %10d微秒\n", name, time.Now().Sub(t).Microseconds())
	for i := 1; i < len(data); i++ {
		if compare(data[i-1], data[i]) > 0 {
			fmt.Println(name, "出错了")
			return
		}
	}
}

func TestSelectSort(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, SelectSort, compare) {
		t.Error("选择排序错误")
	}
}

func TestInsertSort(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, InsertSort, compare) {
		t.Error("插入排序错误")
	}
}

func TestInsertSortList(t *testing.T) {
	l := newList(0, 10000, 10000)
	if !sortListTest(l, InsertSortList, compare) {
		t.Error("链表插入排序错误")
	}
}

func TestBubbleSort(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, BubbleSort, compare) {
		t.Error("冒泡排序错误")
	}
}

func TestBubbleSort2(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, BubbleSort2, compare) {
		t.Error("冒泡排序2错误")
	}
}

func TestBubbleSort3(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, BubbleSort3, compare) {
		t.Error("冒泡排序3错误")
	}
}

func TestShellSort(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, ShellSort, compare) {
		t.Error("希尔排序错误")
	}
}

func TestShellSort2(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, ShellSort2, compare) {
		t.Error("希尔排序2错误")
	}
}

func TestMergeSort(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, MergeSort, compare) {
		t.Error("归并排序错误")
	}
}

func TestMergeSortBU(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, MergeSortBU, compare) {
		t.Error("归并排序2错误")
	}
}

func TestMergeSortList(t *testing.T) {
	l := newList(0, 10000, 10000)
	if !sortListTest(l, MergeSortList, compare) {
		t.Error("链表归并排序错误")
	}
}

func TestQuickSort(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, QuickSort, compare) {
		t.Error("快速排序错误")
	}
}

func TestHeapSort(t *testing.T) {
	data := newInterfaceSlice(0, 10000, 10000)
	if !sortTest(data, HeapSort, compare) {
		t.Error("堆排序错误")
	}
}

// 大概比较一下不同排序算法的效率
func TestSortTime(t *testing.T) {
	n := 1000000
	fmt.Println("测试序列数量级:", n)
	data := make([]int, n)
	// 随机序列
	data = rand.Perm(n)

	// 近乎有序
	//for i := 0; i < n; i++ {
	//	data[i] = i
	//}
	//for k := 0; k < 100; k++ {
	//	i, j := rand.Intn(n), rand.Intn(n)
	//	data[i], data[j] = data[j], data[i]
	//}

	// 大量重复值的序列
	//data = newIntSlice(0, n/4, n)

	//sortTimeTest("选择排序", copyInterfaceSlice(data), SelectSort, compare)
	//sortTimeTest("插入排序", copyInterfaceSlice(data), InsertSort, compare)
	//sortTimeTest("冒泡排序", copyInterfaceSlice(data), BubbleSort, compare)
	//sortTimeTest("冒泡排序2", copyInterfaceSlice(data), BubbleSort2, compare)
	//sortTimeTest("冒泡排序3", copyInterfaceSlice(data), BubbleSort3, compare)
	//sortTimeTest("希尔排序", copyInterfaceSlice(data), ShellSort, compare)
	//sortTimeTest("希尔排序2", copyInterfaceSlice(data), ShellSort2, compare)
	sortTimeTest("归并排序", copyInterfaceSlice(data), MergeSort, compare)
	sortTimeTest("归并排序2", copyInterfaceSlice(data), MergeSortBU, compare)
	sortTimeTest("快速排序", copyInterfaceSlice(data), QuickSort, compare)
	sortTimeTest("堆排序", copyInterfaceSlice(data), HeapSort, compare)
}
