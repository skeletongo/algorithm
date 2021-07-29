package graph

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// TableGraph 邻接表
// 无向，无权图，没有平行边，自环边
type TableGraph struct {
	e    int // 边数量
	data [][]int
}

func (t *TableGraph) SetVertexNumber(n int) {
	t.data = make([][]int, n)
}

func (t *TableGraph) ReadGraph(r io.Reader) error {
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
	t.SetVertexNumber(n)

	line, err = rd.ReadString('\n')
	if err != nil || io.EOF == err {
		return err
	}
	line = strings.TrimRight(line, "\n")
	t.e, err = strconv.Atoi(line)
	if err != nil || io.EOF == err {
		return err
	}

	for {
		line, err = rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimRight(line, "\n")
		vw := strings.Split(line, " ")
		v, err := strconv.Atoi(vw[0])
		if err != nil {
			return err
		}
		w, err := strconv.Atoi(vw[1])
		if err != nil {
			return err
		}
		t.AddEdge(v, w)
	}
	return nil
}

func (t *TableGraph) V() int {
	return len(t.data)
}

func (t *TableGraph) E() int {
	return t.e
}

func (t *TableGraph) AddEdge(v, w int) {
	if v < 0 || w < 0 || v >= len(t.data) || w >= len(t.data) {
		panic("index out of bounds")
	}
	if v == w {
		return
	}
	for _, v := range t.data[v] {
		if v == w {
			return
		}
	}
	t.data[v] = append(t.data[v], w)
	t.data[w] = append(t.data[w], v)
	t.e++
}

func (t *TableGraph) Range(s int, f func(i int, w int) bool) {
	if s < 0 || s >= len(t.data) {
		panic("index out of bounds")
	}
	for k, v := range t.data[s] {
		if !f(k, v) {
			break
		}
	}
}
