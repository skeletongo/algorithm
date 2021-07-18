package search

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func compare(a, b interface{}) int {
	v := a.(float64) - b.(float64)
	if v > 0 {
		return 1
	}
	if v < 0 {
		return -1
	}
	return 0
}

func TestSearchFunc(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var arr []interface{}
		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}
			arr = append(arr, v)
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i].(float64) < arr[j].(float64)
		})

		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}

			s := SearchFunc(arr, compare, v)

			if s == -1 {
				for _, n := range arr {
					if n == v {
						t.Fatal("SearchFunc error")
					}
				}
			} else {
				if arr[s] != v || (s-1 >= 0 && arr[s-1] == v) {
					fmt.Println("SearchFunc error", s, v)
				}
			}
		}
	}
}

func TestFloorFunc(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var arr []interface{}
		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}
			switch rand.Intn(3) {
			case 0:
				v += rand.Float64()
			case 1:
				v -= rand.Float64()
			default:
			}
			arr = append(arr, v)
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i].(float64) < arr[j].(float64)
		})

		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}
			switch rand.Intn(3) {
			case 0:
				v += rand.Float64()
			case 1:
				v -= rand.Float64()
			default:
			}

			s := FloorFunc(arr, compare, v)

			if s == -1 {
				if v > arr[0].(float64) {
					t.Fatal("FloorFunc error -1", arr, v, s)
				}
			} else {
				if arr[s].(float64) >= v || (s+1 < len(arr) && arr[s+1].(float64) < v) {
					fmt.Println("CeilFunc error =", s, v)
				}
			}
		}
	}
}

func TestCeilFunc(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var arr []interface{}
		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}
			switch rand.Intn(3) {
			case 0:
				v += rand.Float64()
			case 1:
				v -= rand.Float64()
			default:
			}
			arr = append(arr, v)
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i].(float64) < arr[j].(float64)
		})

		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}
			switch rand.Intn(3) {
			case 0:
				v += rand.Float64()
			case 1:
				v -= rand.Float64()
			default:
			}

			s := CeilFunc(arr, compare, v)

			if s == -1 {
				if v < arr[len(arr)-1].(float64) {
					t.Fatal("CeilFunc error -1")
				}
			} else {
				if arr[s].(float64) <= v || (s-1 > 0 && arr[s-1].(float64) > v) {
					fmt.Println("CeilFunc error =", s, v)
				}
			}
		}
	}
}

func TestSearch(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var arr []float64
		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}
			arr = append(arr, v)
		}

		sort.Float64s(arr)

		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}

			s := Search(arr, v)

			if s == -1 {
				for _, n := range arr {
					if n == v {
						t.Fatal("Search error")
					}
				}
			} else {
				if arr[s] != v {
					fmt.Println("Search error", s, v)
				}
			}
		}
	}
}

func TestSearchR(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var arr []float64
		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}
			arr = append(arr, v)
		}

		sort.Float64s(arr)

		for j := 0; j < 10; j++ {
			v := float64(rand.Intn(10))
			if rand.Intn(2) == 0 {
				v = -v
			}

			s := SearchR(arr, v)

			if s == -1 {
				for _, n := range arr {
					if n == v {
						t.Fatal("SearchR error")
					}
				}
			} else {
				if arr[s] != v {
					fmt.Println("SearchR error", s, v)
				}
			}
		}
	}
}
