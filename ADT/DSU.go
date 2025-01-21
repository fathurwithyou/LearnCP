package main

import (
	"fmt"
)

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := range dsu.parent {
		dsu.parent[i] = i
	}
	return dsu
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	px, py := d.Find(x), d.Find(y)
	if px == py {
		return
	}
	if d.rank[px] < d.rank[py] {
		px, py = py, px
	}
	d.parent[py] = px
	if d.rank[px] == d.rank[py] {
		d.rank[px]++
	}
}