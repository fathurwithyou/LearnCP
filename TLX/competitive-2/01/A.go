package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Item struct {
	init float64
	v    float64
	cnt  int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].v > pq[j].v }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	pq := &PriorityQueue{}
	heap.Init(pq)

	for i := 0; i < n; i++ {
		var v float64
		fmt.Fscanf(reader, "%f", &v)
		heap.Push(pq, Item{init: v, v: v, cnt: 1})
	}

	m -= n

	for m > 0 {
		top := heap.Pop(pq).(Item)
		newCnt := top.cnt + 1
		newValue := (top.init / float64(newCnt))

		heap.Push(pq, Item{init: top.init, v: newValue, cnt: newCnt})
		m--
	}

	top := heap.Pop(pq).(Item)
	fmt.Fprintf(writer, "%.2f", float64(top.v))
}
