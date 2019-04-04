package algorithm

import (
	"math/rand"
	"testing"
	"time"
)

var arr []int

// 生成测试数据
func testdata() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr = arr[:0]
	for i := 0; i < 100; i++ {
		arr = append(arr, r.Intn(1000))
	}
	//fmt.Println(arr)
}

// 排序检查
func check() bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func TestSelectSort(t *testing.T) {
	testdata()
	SelectSort(arr)
	if !check() {
		t.Error("选择排序错误")
	}
}

func TestInsertSort(t *testing.T) {
	testdata()
	InsertSort(arr)
	if !check() {
		t.Error("插入排序错误")
	}
}

func TestBubbleSort(t *testing.T) {
	testdata()
	BubbleSort(arr)
	if !check() {
		t.Error("冒泡排序错误")
	}
	testdata()
	BubbleSort2(arr)
	if !check() {
		t.Error("冒泡排序错误2")
	}
	testdata()
	BubbleSort3(arr)
	if !check() {
		t.Error("冒泡排序错误3")
	}
}

func TestShellSort(t *testing.T) {
	testdata()
	ShellSort(arr)
	if !check() {
		t.Error("希尔排序错误")
	}
	testdata()
	ShellSort2(arr)
	if !check() {
		t.Error("希尔排序错误2")
	}
}
