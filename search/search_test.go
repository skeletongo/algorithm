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

func TestFloor(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var arr []float64
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

		sort.Float64s(arr)

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

			s := Floor(arr, v)

			if s == -1 {
				if v >= arr[0] {
					t.Fatal("Floor error -1")
				}
			} else {
				if arr[s] == v && s-1 > 0 && arr[s-1] == v {
					fmt.Println("Floor error =", s, v)
				}
				if arr[s] != v && (arr[s] > v || (s+1 < len(arr) && arr[s+1] < v)) {
					fmt.Println("Floor error !=", s, v)
				}
			}
		}
	}
}

func TestCeil(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var arr []float64
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

		sort.Float64s(arr)

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

			s := Ceil(arr, v)

			if s == -1 {
				if v <= arr[0] {
					t.Fatal("Ceil error -1")
				}
			} else {
				if arr[s] == v && s+1 < len(arr) && arr[s+1] == v {
					fmt.Println("Ceil error =", s, v)
				}
				if arr[s] != v && (arr[s] < v || (s-1 > 0 && arr[s-1] > v)) {
					fmt.Println("Ceil error !=", s, v)
				}
			}
		}
	}
}
