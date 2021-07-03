package search

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSearch(t *testing.T) {
	n := 10000
	var arr []float64
	for i := 0; i < 10; i++ {
		arr = append(arr, float64(i))
	}

	for i := 0; i < n; i++ {
		v := float64(rand.Intn(n))
		if rand.Intn(2) == 0 {
			v = -v
		}
		if rand.Intn(2) == 0 {
			v += rand.Float64()
		} else {
			v -= rand.Float64()
		}
		s := Search(arr, v)

		//fmt.Println(v,s)

		if s == -1 {
			if v >= 0 && v < 10 && float64(int(v)) == v {
				fmt.Println("Search error", s, v)
			}
		} else {
			if arr[s] != v {
				fmt.Println("Search error", s, v)
			}
		}
	}
}

func TestSearch2(t *testing.T) {
	n := 10000
	var arr []float64
	for i := 0; i < 10; i++ {
		arr = append(arr, float64(i))
	}

	for i := 0; i < n; i++ {
		v := float64(rand.Intn(n))
		if rand.Intn(2) == 0 {
			v = -v
		}
		if rand.Intn(2) == 0 {
			v += rand.Float64()
		} else {
			v -= rand.Float64()
		}
		s := Search2(arr, v)

		//fmt.Println(v,s)

		if s == -1 {
			if v >= 0 && v < 10 && float64(int(v)) == v {
				fmt.Println("Search error", s, v)
			}
		} else {
			if arr[s] != v {
				fmt.Println("Search error", s, v)
			}
		}
	}
}

func TestFloor(t *testing.T) {
	n := 10000
	var arr []float64
	for i := 0; i < 10; i++ {
		arr = append(arr, float64(i))
	}

	for i := 0; i < n; i++ {
		v := float64(rand.Intn(n))
		if rand.Intn(2) == 0 {
			v = -v
		}
		if rand.Intn(2) == 0 {
			v += rand.Float64()
		} else {
			v -= rand.Float64()
		}
		s := Floor(arr, v)

		//fmt.Println(v,s)

		if s == -1 {
			if v >= 0 {
				fmt.Println("Floor error", s, v)
			}
		} else {
			if arr[s] > v {
				fmt.Println("Floor error", s, v)
			}
		}
	}
}

func TestCeil(t *testing.T) {
	n := 10000
	var arr []float64
	for i := 0; i < 10; i++ {
		arr = append(arr, float64(i))
	}

	for i := 0; i < n; i++ {
		v := float64(rand.Intn(n))
		if rand.Intn(2) == 0 {
			v = -v
		}
		if rand.Intn(2) == 0 {
			v += rand.Float64()
		} else {
			v -= rand.Float64()
		}
		s := Ceil(arr, v)

		//fmt.Println(v,s)

		if s == -1 {
			if v < 9 {
				fmt.Println("Ceil error", s, v)
			}
		} else {
			if arr[s] < v {
				fmt.Println("Ceil error", s, v)
			}
		}
	}
}
