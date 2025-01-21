package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	v, w int
}

type Item struct {
	val interface{}
}

type Queue []Item

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(val interface{}) {
	*q = append(*q, Item{val: val})
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func (q *Queue) Front() interface{} {
	if len(*q) == 0 {
		return nil
	}
	return (*q)[0].val
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Back() interface{} {
	if len(*q) == 0 {
		return nil
	}
	return (*q)[len(*q)-1].val
}

func (q *Queue) Pop() {
	if !(*q).Empty() {
		*q = (*q)[1:]
	}
}

func dijkstra(adj [][]Pair, src int) []int {
	dist := make([]int, len(adj))
	for i := range dist {
		dist[i] = 1e9
	}
	dist[src] = 0
	q := NewQueue()
	q.Push(src)
	for !q.Empty() {
		u := q.Front().(int)
		q.Pop()
		for _, p := range adj[u] {
			v, w := p.v, p.w
			if dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
				q.Push(v)
			}
		}
	}
	return dist
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	params := strings.Fields(line)
	n, _ := strconv.Atoi(params[0])
	m, _ := strconv.Atoi(params[1])
	p, _ := strconv.Atoi(params[2])
	q, _ := strconv.Atoi(params[3])

	adj := make([][]Pair, n+1)
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		edge := strings.Fields(line)
		u, _ := strconv.Atoi(edge[0])
		v, _ := strconv.Atoi(edge[1])
		w, _ := strconv.Atoi(edge[2])
		adj[u] = append(adj[u], Pair{v: v, w: w})
		adj[v] = append(adj[v], Pair{v: u, w: w})
	}

	dist := dijkstra(adj, p)
	fmt.Fprintln(writer, dist[q])
}