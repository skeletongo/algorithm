package graph

import (
	"io"
)

// Digraph 有向图
type Digraph interface {
	SetVertexNumber(n int)
	ReadGraph(r io.Reader) error
	V() int
	E() int
	AddEdge(v, w int)
	Range(s int, f func(i, w int) bool)
	Revers() Digraph
}

type DirectedDFS struct {
	marked []bool
}

// NewDirectedDFS 查找s可以到达的所有顶点
func NewDirectedDFS(g Digraph, s int) *DirectedDFS {
	if s < 0 || s >= g.V() {
		panic("index out of bounds")
	}
	d := new(DirectedDFS)
	d.marked = make([]bool, g.V())
	d.dfs(g, s)
	return d
}

// NewDirectedDFSSource 查找source可以到达的所有顶点
func NewDirectedDFSSource(g Digraph, source []int) *DirectedDFS {
	d := new(DirectedDFS)
	d.marked = make([]bool, g.V())
	for _, v := range source {
		if v < 0 || v >= g.V() {
			panic("index out of bounds")
		}
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}
	return d
}

func (d *DirectedDFS) dfs(g Graph, v int) {
	d.marked[v] = true
	g.Range(v, func(i, w int) bool {
		if !d.marked[w] {
			d.dfs(g, w)
		}
		return true
	})
}

// Marked v是否可达
func (d *DirectedDFS) Marked(v int) bool {
	if v < 0 || v >= len(d.marked) {
		panic("index out of bounds")
	}
	return d.marked[v]
}

type DirectedCycle struct {
	marked []bool // 标记顶点是否检查过
	edgeTo []int  // 有向边，索引为起始顶点，值为到达顶点
	stack  []bool // 在栈中的顶部标记为true
	cycle  []int
}

// NewDirectedCycle 寻找有向环
// 环：至少两条边两个顶点
func NewDirectedCycle(g Digraph) *DirectedCycle {
	d := new(DirectedCycle)
	d.marked = make([]bool, g.V())
	d.edgeTo = make([]int, g.V())
	d.stack = make([]bool, g.V())
	for i := 0; i < g.V(); i++ {
		if !d.marked[i] {
			d.dfs(g, i)
		}
	}
	return d
}

func (d *DirectedCycle) dfs(g Digraph, v int) {
	d.marked[v] = true
	// 开始访问顶点入栈
	d.stack[v] = true
	g.Range(v, func(i, w int) bool {
		if d.HasCycle() {
			return false
		}
		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		} else if d.stack[w] {
			d.cycle = make([]int, 0, 3)
			for x := v; x != w; x = d.edgeTo[x] {
				d.cycle = append(d.cycle, x)
			}
			d.cycle = append(d.cycle, w)
			d.cycle = append(d.cycle, v)
			return false
		}
		return true
	})
	// 顶点访问结束出站
	d.stack[v] = false
}

// HasCycle 是否有环
func (d *DirectedCycle) HasCycle() bool {
	return d.cycle != nil
}

// Cycle 返回找到的第一个环
func (d *DirectedCycle) Cycle() []int {
	return d.cycle
}

type DepthFirstOrder struct {
	marked                 []bool
	pre, post, reversePost []int
}

// NewDepthFirstOrder 有向图基于深度优先搜索排序
func NewDepthFirstOrder(g Digraph) *DepthFirstOrder {
	d := new(DepthFirstOrder)
	d.marked = make([]bool, g.V())
	d.pre = make([]int, 0, g.V())
	d.post = make([]int, 0, g.V())
	d.reversePost = make([]int, 0, g.V())
	for i := 0; i < g.V(); i++ {
		if !d.marked[i] {
			d.dfs(g, i)
		}
	}
	for i := g.V() - 1; i >= 0; i-- {
		d.reversePost = append(d.reversePost, d.post[i])
	}
	return d
}

func (d *DepthFirstOrder) dfs(g Digraph, v int) {
	d.marked[v] = true
	d.pre = append(d.pre, v)
	g.Range(v, func(i, w int) bool {
		if !d.marked[w] {
			d.dfs(g, w)
		}
		return true
	})
	d.post = append(d.post, v)
}

// Pre 前序排列
func (d *DepthFirstOrder) Pre() []int {
	return d.pre
}

// Post 后序排列
func (d *DepthFirstOrder) Post() []int {
	return d.post
}

// ReversePost 逆后序排列
func (d *DepthFirstOrder) ReversePost() []int {
	return d.reversePost
}

type TopLogical struct {
	order []int
}

// NewTopLogical 对有向无环图进行拓扑排序
func NewTopLogical(g Digraph) *TopLogical {
	t := new(TopLogical)
	// 检测是否有环，有环图不能拓扑排序
	c := NewDirectedCycle(g)
	if c.HasCycle() {
		return t
	}
	d := NewDepthFirstOrder(g)
	t.order = d.ReversePost()
	return t
}

// IsDAG 是不是有向无环图
func (t *TopLogical) IsDAG() bool {
	return t.order != nil
}

// Order 拓扑排列
func (t *TopLogical) Order() []int {
	return t.order
}

type SCC struct {
	n      int
	id     []int
	marked []bool
}

// NewSCC 查询有向图的强连通分量
// 方法: kosaraju算法
// 先求有向图的反向图然后再求反向图的连通分量就是原有向图的强连通分量
func NewSCC(g Digraph) *SCC {
	s := new(SCC)
	s.id = make([]int, g.V())
	s.marked = make([]bool, g.V())
	d := NewDepthFirstOrder(g.Revers())
	for _, v := range d.ReversePost() {
		if !s.marked[v] {
			s.dfs(g, v)
			s.n++
		}
	}
	return s
}

func (s *SCC) dfs(g Digraph, v int) {
	s.marked[v] = true
	s.id[v] = s.n
	g.Range(v, func(i, w int) bool {
		if !s.marked[w] {
			s.dfs(g, w)
		}
		return true
	})
}

// StronglyConnected v和w是否强连通
func (s *SCC) StronglyConnected(v, w int) bool {
	return s.Id(v) == s.Id(w)
}

// Count 强连通分量数量
func (s *SCC) Count() int {
	return s.n
}

// Id v所在的强连通分量值
func (s *SCC) Id(v int) int {
	if v < 0 || v >= len(s.marked) {
		panic("index out of bounds")
	}
	return s.id[v]
}

type TransitiveClosure struct {
	data []*DirectedDFS
}

// TransitiveClosure 查询顶点对的可达性
func NewTransitiveClosure(g Digraph) *TransitiveClosure {
	t := new(TransitiveClosure)
	t.data = make([]*DirectedDFS, g.V())
	for i := 0; i < g.V(); i++ {
		t.data[i] = NewDirectedDFS(g, i)
	}
	return t
}

// Reachable 顶点w是否从顶点v可达
func (t *TransitiveClosure) Reachable(v, w int) bool {
	if v < 0 || v >= len(t.data) {
		panic("index out of bounds")
	}
	return t.data[v].Marked(w)
}
