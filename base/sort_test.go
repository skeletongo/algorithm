package base

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newIntSlice(l, r int, n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = l + rand.Intn(r-l+1)
	}
	return ret
}

func sortFuncTest(data []int, f func(arr []int, l, r int)) bool {
	f(data, 0, len(data)-1)
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func sortTimeTest(name string, data []int, f func(arr []int, l, r int)) {
	t := time.Now()
	f(data, 0, len(data)-1)
	fmt.Printf("%8s: %10d微秒\n", name, time.Now().Sub(t).Microseconds())
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			fmt.Println(name, "出错了")
			return
		}
	}
}

func TestSelectSort(t *testing.T) {
	data := newIntSlice(0, 10000, 10000)
	if !sortFuncTest(data, SelectSort) {
		t.Error("选择排序错误")
	}
}

func TestInsertSort(t *testing.T) {
	data := newIntSlice(0, 10000, 10000)
	if !sortFuncTest(data, InsertSort) {
		t.Error("插入排序错误")
	}
}

func TestBubbleSort(t *testing.T) {
	data := newIntSlice(0, 10000, 10000)
	if !sortFuncTest(data, BubbleSort) {
		t.Error("冒泡排序错误")
	}
}

func TestBubbleSort2(t *testing.T) {
	data := newIntSlice(0, 10000, 10000)
	if !sortFuncTest(data, BubbleSort2) {
		t.Error("冒泡排序2错误")
	}
}

func TestBubbleSort3(t *testing.T) {
	data := newIntSlice(0, 10000, 10000)
	if !sortFuncTest(data, BubbleSort3) {
		t.Error("冒泡排序3错误")
	}
}

func TestShellSort(t *testing.T) {
	data := newIntSlice(0, 10000, 10000)
	if !sortFuncTest(data, ShellSort) {
		t.Error("希尔排序错误")
	}
}

func TestShellSort2(t *testing.T) {
	data := newIntSlice(0, 10000, 10000)
	if !sortFuncTest(data, ShellSort2) {
		t.Error("希尔排序2错误")
	}
}

func TestMergeSort(t *testing.T) {
	data := newIntSlice(0, 10000, 10000)
	if !sortFuncTest(data, MergeSort) {
		t.Error("归并排序错误")
	}
}

func TestMergeSortBU(t *testing.T) {
	data := newIntSlice(0, 10000, 10000)
	if !sortFuncTest(data, MergeSortBU) {
		t.Error("归并排序2错误")
	}
}

func copyIntSlice(data []int) []int {
	ret := make([]int, len(data))
	copy(ret, data)
	return ret
}

// 大概比较一下不同排序算法的效率
func TestSortTime(t *testing.T) {
	n := 10000
	data := rand.Perm(n)
	fmt.Println("测试序列数量级:", n)
	sortTimeTest("选择排序", copyIntSlice(data), SelectSort)

	sortTimeTest("插入排序", copyIntSlice(data), InsertSort)

	sortTimeTest("冒泡排序", copyIntSlice(data), BubbleSort)
	sortTimeTest("冒泡排序2", copyIntSlice(data), BubbleSort2)
	sortTimeTest("冒泡排序3", copyIntSlice(data), BubbleSort3)

	sortTimeTest("希尔排序", copyIntSlice(data), ShellSort)
	sortTimeTest("希尔排序2", copyIntSlice(data), ShellSort2)

	sortTimeTest("归并排序", copyIntSlice(data), MergeSort)
	sortTimeTest("归并排序2", copyIntSlice(data), MergeSortBU)
}
