package graph

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// MatrixGraph 邻接矩阵
// 无向，无权图，没有平行边，自环边
type MatrixGraph struct {
	v    int // 节点数量
	e    int // 边数量
	data [][]bool
}

func (m *MatrixGraph) SetVertexNumber(n int) {
	for i := 0; i < n; i++ {
		m.data = append(m.data, make([]bool, n))
	}
	m.v = n
}

func (m *MatrixGraph) ReadGraph(r io.Reader) error {
	rd := bufio.NewReader(r)
	line, err := rd.ReadString('\n')
	if err != nil || io.EOF == err {
		return err
	}
	line = strings.TrimRight(line, "\n")
	n, err := strconv.Atoi(line)
	if err != nil {
		return err
	}
	m.SetVertexNumber(n)

	line, err = rd.ReadString('\n')
	if err != nil || io.EOF == err {
		return err
	}
	line = strings.TrimRight(line, "\n")
	e, err := strconv.Atoi(line)
	if err != nil {
		return err
	}

	for i := 0; i < e; i++ {
		line, _ = rd.ReadString('\n')
		line = strings.TrimRight(line, "\n")
		if line == "" {
			return errors.New("edge error")
		}
		vw := strings.Split(line, " ")
		v, err := strconv.Atoi(vw[0])
		if err != nil {
			return err
		}
		w, err := strconv.Atoi(vw[1])
		if err != nil {
			return err
		}
		m.AddEdge(v, w)
	}
	return nil
}

func (m *MatrixGraph) V() int {
	return m.v
}

func (m *MatrixGraph) E() int {
	return m.e
}

func (m *MatrixGraph) AddEdge(v, w int) {
	if v < 0 || w < 0 || v >= m.v || w >= m.v {
		panic("index out of bounds")
	}
	if v == w {
		return
	}
	if m.data[v][w] {
		return
	}
	m.data[v][w] = true
	m.data[w][v] = true
	m.e++
}

func (m *MatrixGraph) Range(s int, f func(i int, w int) bool) {
	if s < 0 || s >= m.v {
		panic("index out of bounds")
	}
	i := 0
	for k, v := range m.data[s] {
		if v {
			if !f(i, k) {
				break
			}
			i++
		}
	}
}

func (m *MatrixGraph) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("顶点数: %v\n", m.v))
	buf.WriteString(fmt.Sprintf("边数: %v\n", m.e))
	n := len(fmt.Sprint(m.v)) + 2
	buf.WriteString(strings.Join(make([]string, n+1), " "))
	s := "%-" + fmt.Sprint(n) + "v"
	for i := 0; i < m.v; i++ {
		buf.WriteString(fmt.Sprintf(s, i))
	}
	buf.WriteString("\n")
	for i := 0; i < len(m.data); i++ {
		buf.WriteString(fmt.Sprintf(s, i))
		for j := 0; j < len(m.data[i]); j++ {
			if m.data[i][j] {
				buf.WriteString(fmt.Sprintf(s, 1))
			} else {
				buf.WriteString(fmt.Sprintf(s, 0))
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

// MatrixDigraph 邻接矩阵
// 有向，无权图，没有平行边，自环边
type MatrixDigraph struct {
	v    int // 节点数量
	e    int // 边数量
	data [][]bool
}

func (m *MatrixDigraph) SetVertexNumber(n int) {
	for i := 0; i < n; i++ {
		m.data = append(m.data, make([]bool, n))
	}
	m.v = n
}

func (m *MatrixDigraph) ReadGraph(r io.Reader) error {
	rd := bufio.NewReader(r)
	line, err := rd.ReadString('\n')
	if err != nil || io.EOF == err {
		return err
	}
	line = strings.TrimRight(line, "\n")
	n, err := strconv.Atoi(line)
	if err != nil {
		return err
	}
	m.SetVertexNumber(n)

	line, err = rd.ReadString('\n')
	if err != nil || io.EOF == err {
		return err
	}
	line = strings.TrimRight(line, "\n")
	e, err := strconv.Atoi(line)
	if err != nil {
		return err
	}

	for i := 0; i < e; i++ {
		line, _ = rd.ReadString('\n')
		line = strings.TrimRight(line, "\n")
		if line == "" {
			return errors.New("edge error")
		}
		vw := strings.Split(line, " ")
		v, err := strconv.Atoi(vw[0])
		if err != nil {
			return err
		}
		w, err := strconv.Atoi(vw[1])
		if err != nil {
			return err
		}
		m.AddEdge(v, w)
	}
	return nil
}

func (m *MatrixDigraph) V() int {
	return m.v
}

func (m *MatrixDigraph) E() int {
	return m.e
}

func (m *MatrixDigraph) AddEdge(v, w int) {
	if v < 0 || w < 0 || v >= m.v || w >= m.v {
		panic("index out of bounds")
	}
	if v == w {
		return
	}
	if m.data[v][w] {
		return
	}
	m.data[v][w] = true
	m.e++
}

func (m *MatrixDigraph) Range(s int, f func(i int, w int) bool) {
	if s < 0 || s >= m.v {
		panic("index out of bounds")
	}
	i := 0
	for k, v := range m.data[s] {
		if v {
			if !f(i, k) {
				break
			}
			i++
		}
	}
}

func (m *MatrixDigraph) Revers() Digraph {
	g := new(MatrixDigraph)
	g.SetVertexNumber(m.v)
	g.e = m.e
	for i := 0; i < len(m.data); i++ {
		for j := 0; j < len(m.data[i]); j++ {
			if m.data[i][j] {
				g.data[j][i] = true
			}
		}
	}
	return g
}

func (m *MatrixDigraph) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("顶点数: %v\n", m.v))
	buf.WriteString(fmt.Sprintf("边数: %v\n", m.e))
	n := len(fmt.Sprint(m.v)) + 2
	buf.WriteString(strings.Join(make([]string, n+1), " "))
	s := "%-" + fmt.Sprint(n) + "v"
	for i := 0; i < m.v; i++ {
		buf.WriteString(fmt.Sprintf(s, i))
	}
	buf.WriteString("\n")
	for i := 0; i < len(m.data); i++ {
		buf.WriteString(fmt.Sprintf(s, i))
		for j := 0; j < len(m.data[i]); j++ {
			if m.data[i][j] {
				buf.WriteString(fmt.Sprintf(s, 1))
			} else {
				buf.WriteString(fmt.Sprintf(s, 0))
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}
