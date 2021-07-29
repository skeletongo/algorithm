package graph

import (
	"container/list"
	"io"
)

// Graph 无向图
type Graph interface {
	SetVertexNumber(n int)
	ReadGraph(r io.Reader) error
	V() int
	E() int
	AddEdge(v, w int)
	Range(s int, f func(i, w int) bool)
}

type Search struct {
	marked []bool
	count  int
}

// NewSearch 找出和s连通的所有顶点
func NewSearch(g Graph, s int) *Search {
	if s < 0 || s >= g.V() {
		panic("index out of bounds")
	}
	search := new(Search)
	search.marked = make([]bool, g.V())
	search.dfs(g, s)
	return search
}

// dfs 深度优先遍历
func (s *Search) dfs(g Graph, v int) {
	s.marked[v] = true
	s.count++
	g.Range(v, func(i, w int) bool {
		if !s.marked[w] {
			s.dfs(g, w)
		}
		return true
	})
}

// Marked v和s是否连通
func (s *Search) Marked(v int) bool {
	if v < 0 || v >= len(s.marked) {
		panic("index out of bounds")
	}
	return s.marked[v]
}

// Count 和s连通的顶点数量
func (s *Search) Count() int {
	return s.count
}

type Paths struct {
	marked []bool
	edgeTo []int
	s      int
}

// NewPaths 找出任意顶点到s起点的最短路径
func NewPaths(g Graph, s int) *Paths {
	if s < 0 || s >= g.V() {
		panic("index out of bounds")
	}
	paths := new(Paths)
	paths.marked = make([]bool, g.V())
	paths.edgeTo = make([]int, g.V())
	paths.s = s
	paths.bfs(g, s)
	return paths
}

// bfs 广度优先遍历
func (p *Paths) bfs(g Graph, v int) {
	q := list.New()
	q.PushBack(v)
	p.marked[v] = true
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		v = e.Value.(int)
		g.Range(v, func(i, w int) bool {
			if !p.marked[w] {
				p.marked[w] = true
				p.edgeTo[w] = v
				q.PushBack(w)
			}
			return true
		})
	}
}

// HasPathTo 是否有s到v的最短路径
func (p *Paths) HasPathTo(v int) bool {
	if v < 0 || v >= len(p.marked) {
		panic("index out of bounds")
	}
	return p.marked[v]
}

// PathTo 找出s到v的最短路径
func (p *Paths) PathTo(v int) []int {
	if !p.HasPathTo(v) {
		return nil
	}
	var ret []int
	for x := v; x != p.s; x = p.edgeTo[x] {
		ret = append(ret, x)
	}
	ret = append(ret, p.s)
	return ret
}

type CC struct {
	n      int
	id     []int
	marked []bool
}

// NewCC 计算图的连通分量
func NewCC(g Graph) *CC {
	cc := new(CC)
	cc.id = make([]int, g.V())
	cc.marked = make([]bool, g.V())
	for i := 0; i < g.V(); i++ {
		if !cc.marked[i] {
			cc.dfs(g, i)
			cc.n++
		}
	}
	return cc
}

func (c *CC) dfs(g Graph, v int) {
	c.marked[v] = true
	c.id[v] = c.n
	g.Range(v, func(i, w int) bool {
		if !c.marked[w] {
			c.dfs(g, w)
		}
		return true
	})
}

// Connected 两顶点是否连通
func (c *CC) Connected(v, w int) bool {
	return c.Id(v) == c.Id(w)
}

// Count 连通分量数量
func (c *CC) Count() int {
	return c.n
}

// Id 查询顶点所在的连通分量
func (c *CC) Id(v int) int {
	if v < 0 || v >= len(c.marked) {
		panic("index out of bounds")
	}
	return c.id[v]
}

type Cycle struct {
	marked []bool
	has    bool
}

// NewCycle 判断是否有环
// 环：最少三个节点
func NewCycle(g Graph) *Cycle {
	c := new(Cycle)
	c.marked = make([]bool, g.V())
	for i := 0; i < g.V(); i++ {
		if !c.marked[i] {
			c.hasCycle(g, i, i)
		}
	}
	return c
}

func (c *Cycle) hasCycle(g Graph, v, p int) bool {
	c.marked[v] = true
	g.Range(v, func(i, w int) bool {
		if !c.marked[w] {
			return c.hasCycle(g, w, v)
		} else if w != p {
			c.has = true
			return false
		}
		return true
	})
	return true
}

// HasCycle 是否有环
func (c *Cycle) HasCycle() bool {
	return c.has
}

type BSTGraph struct {
	marked []bool
	color  []bool
	isBst  bool
}

// NewBSTGraph 判断是否是二分图
func NewBSTGraph(g Graph) *BSTGraph {
	b := new(BSTGraph)
	b.marked = make([]bool, g.V())
	b.color = make([]bool, g.V())
	b.isBst = true
	for i := 0; i < g.V(); i++ {
		if !b.marked[i] {
			b.isBSTGraph(g, i)
		}
	}
	return b
}

func (b *BSTGraph) isBSTGraph(g Graph, v int) bool {
	b.marked[v] = true
	g.Range(v, func(i, w int) bool {
		if !b.marked[w] {
			b.color[w] = !b.color[v]
			b.isBSTGraph(g, w)
		} else if b.color[w] == b.color[v] {
			b.isBst = false
			return false
		}
		return true
	})
	return true
}

// IsBSTGraph 是不是二分图
func (b *BSTGraph) IsBSTGraph() bool {
	return b.isBst
}
