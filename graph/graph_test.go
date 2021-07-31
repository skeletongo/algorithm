package graph

import (
	"fmt"
	"io"
	"os"
	"testing"
)

var matrix *MatrixGraph
var table *TableGraph
var dimatrix *MatrixDigraph
var ditable *TableDigraph
var dimatrixNoLoop *MatrixDigraph
var ditableNoLoop *TableDigraph

func init() {
	matrix = new(MatrixGraph)
	if err := matrix.ReadGraph(readTestFile("graph.txt")); err != nil {
		panic(err)
	}
	fmt.Println(matrix)

	table = new(TableGraph)
	if err := table.ReadGraph(readTestFile("graph.txt")); err != nil {
		panic(err)
	}
	fmt.Println(table)

	dimatrix = new(MatrixDigraph)
	if err := dimatrix.ReadGraph(readTestFile("digraph.txt")); err != nil {
		panic(err)
	}
	fmt.Println(dimatrix)

	ditable = new(TableDigraph)
	if err := ditable.ReadGraph(readTestFile("digraph.txt")); err != nil {
		panic(err)
	}
	fmt.Println(ditable)

	dimatrixNoLoop = new(MatrixDigraph)
	if err := dimatrixNoLoop.ReadGraph(readTestFile("digraph_no_loop.txt")); err != nil {
		panic(err)
	}
	fmt.Println(dimatrixNoLoop)

	ditableNoLoop = new(TableDigraph)
	if err := ditableNoLoop.ReadGraph(readTestFile("digraph_no_loop.txt")); err != nil {
		panic(err)
	}
	fmt.Println(ditableNoLoop)
}

func readTestFile(filename string) io.ReadCloser {
	f, err := os.Open(filename)
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

func TestNewDirectedDFS(t *testing.T) {
	d := NewDirectedDFS(dimatrix, 0)
	for i := 0; i < dimatrix.V(); i++ {
		fmt.Printf("0->%d %v\n", i, d.Marked(i))
	}
	fmt.Println()

	d = NewDirectedDFSSource(dimatrix, []int{0, 6, 7})
	for i := 0; i < dimatrix.V(); i++ {
		fmt.Printf("0->%d %v\n", i, d.Marked(i))
	}
	fmt.Println()

	d = NewDirectedDFS(ditable, 0)
	for i := 0; i < ditable.V(); i++ {
		fmt.Printf("0->%d %v\n", i, d.Marked(i))
	}
	fmt.Println()

	d = NewDirectedDFSSource(ditable, []int{0, 6, 7})
	for i := 0; i < ditable.V(); i++ {
		fmt.Printf("0->%d %v\n", i, d.Marked(i))
	}
	fmt.Println()
}

func TestNewDirectedCycle(t *testing.T) {
	d := NewDirectedCycle(dimatrix)
	fmt.Printf("%v %v\n", d.HasCycle(), d.Cycle())
	d = NewDirectedCycle(ditable)
	fmt.Printf("%v %v\n", d.HasCycle(), d.Cycle())
}

func TestNewDepthFirstOrder(t *testing.T) {
	d := NewDepthFirstOrder(dimatrixNoLoop)
	fmt.Println("前序排列:", d.Pre())
	fmt.Println("后序排列:", d.Post())
	fmt.Println("逆后序排列:", d.ReversePost())

	d = NewDepthFirstOrder(ditableNoLoop)
	fmt.Println("前序排列:", d.Pre())
	fmt.Println("后序排列:", d.Post())
	fmt.Println("逆后序排列:", d.ReversePost())
}

func TestNewTopLogical(t *testing.T) {
	top := NewTopLogical(dimatrixNoLoop)
	fmt.Println(top.IsDAG(), top.Order())

	top = NewTopLogical(ditableNoLoop)
	fmt.Println(top.IsDAG(), top.Order())
}

func TestNewSCC(t *testing.T) {
	s := NewSCC(dimatrix)
	fmt.Println("强连通分量数:", s.Count())
	for i := 0; i < ditable.V(); i++ {
		fmt.Printf("%v 强连通分量值:%v\n", i, s.Id(i))
	}
	for i := 0; i < ditable.V(); i++ {
		fmt.Printf("0-%d 强连通:%v\n", i, s.StronglyConnected(0, i))
	}

	s = NewSCC(ditable)
	fmt.Println("强连通分量数:", s.Count())
	for i := 0; i < ditable.V(); i++ {
		fmt.Printf("%v 强连通分量值:%v\n", i, s.Id(i))
	}
	for i := 0; i < ditable.V(); i++ {
		fmt.Printf("0-%d 强连通:%v\n", i, s.StronglyConnected(0, i))
	}
}

func TestNewTransitiveClosure(t *testing.T) {
	tc := NewTransitiveClosure(dimatrix)
	for i := 0; i < dimatrix.V(); i++ {
		fmt.Printf("0->%d %v\n", i, tc.Reachable(0, i))
	}

	tc = NewTransitiveClosure(ditable)
	for i := 0; i < dimatrix.V(); i++ {
		fmt.Printf("0->%d %v\n", i, tc.Reachable(0, i))
	}
}
