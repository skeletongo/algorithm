package graph

import (
	"fmt"
	"io"
	"os"
	"testing"
)

var matrix *MatrixGraph
var table *TableGraph

func init() {
	matrix = new(MatrixGraph)
	if err := matrix.ReadGraph(readTestFile()); err != nil {
		panic(err)
	}
	table = new(TableGraph)
	if err := table.ReadGraph(readTestFile()); err != nil {
		panic(err)
	}
}

func readTestFile() io.ReadCloser {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	return f
}

func TestNewSearch(t *testing.T) {
	s := NewSearch(matrix, 0)
	for i := 0; i < matrix.V(); i++ {
		fmt.Printf("0 -> %d %v\n", i, s.Marked(i))
	}
	fmt.Println(s.Count())

	fmt.Println()

	s1 := NewSearch(table, 0)
	for i := 0; i < table.V(); i++ {
		fmt.Printf("0 -> %d %v\n", i, s1.Marked(i))
	}
	fmt.Println(s1.Count())
}

func TestNewPaths(t *testing.T) {
	p := NewPaths(matrix, 0)
	for i := 0; i < matrix.V(); i++ {
		fmt.Println(p.PathTo(i))
	}

	fmt.Println()

	p1 := NewPaths(table, 0)
	for i := 0; i < table.V(); i++ {
		fmt.Println(p1.PathTo(i))
	}
}

func TestNewCC(t *testing.T) {
	c := NewCC(matrix)
	for i := 0; i < matrix.V(); i++ {
		fmt.Printf("%d %d %v\n", 0, i, c.Connected(0, i))
	}
	fmt.Println(c.Count())

	fmt.Println()

	c1 := NewCC(table)
	for i := 0; i < table.V(); i++ {
		fmt.Printf("%d %d %v\n", 0, i, c.Connected(0, i))
	}
	fmt.Println(c1.Count())
}

func TestNewCycle(t *testing.T) {
	c := NewCycle(matrix)
	fmt.Println(c.HasCycle())

	c = NewCycle(table)
	fmt.Println(c.HasCycle())
}

func TestNewBSTGraph(t *testing.T) {
	b := NewBSTGraph(matrix)
	fmt.Println(b.IsBSTGraph())

	b = NewBSTGraph(table)
	fmt.Println(b.IsBSTGraph())
}
