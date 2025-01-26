package main

import (
	"fmt"
) 

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

var vis [1001][1001]bool
var arr [1001]string

var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}
var n, m int

var sx, sy, tx, ty int

func main(){

	fmt.Scan(&n, &m)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if arr[i][j] == 'A' {
				sx = i
				sy = j
			} else if arr[i][j] == 'B' {
				tx = i
				ty = j
			}
		}
	}

	q := NewQueue()
	q.Push([]int{sx, sy})
	vis[sx][sy] = true
	dist[sx][sy] = 0
	found := false
	
	for !q.Empty() {
		front := q.Front().([]int)
		q.Pop()
		x := front[0]
		y := front[1]
		if x == tx && y == ty {	
			found = true
			break
		}

		for k := 0; k < 4; k++ {
			nx := x + dx[k]
			ny := y + dy[k]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && arr[nx][ny] != '#' && !vis[nx][ny] {
				vis[nx][ny] = true
				dist[nx][ny] = dist[x][y] + 1
				q.Push([]int{nx, ny})
			}
		}
	}

	if !found {
		fmt.Println("NO")
		return
	}else{
		fmt.Println("YES")
		fmt.Println(dist[tx][ty])
		ans := make([]byte, dist[tx][ty])
		for x, y := tx, ty; x != sx || y != sy; {
			for k := 0; k < 4; k++ {
				nx := x + dx[k]
				ny := y + dy[k]
				if nx >= 0 && nx < n && ny >= 0 && ny < m && arr[nx][ny] != '#' && dist[nx][ny] == dist[x][y] - 1 {
					if k == 0 {
						ans[dist[x][y]-1] = 'L'
					} else if k == 1 {
						ans[dist[x][y]-1] = 'R'
					} else if k == 2 {
						ans[dist[x][y]-1] = 'U'
					} else {
						ans[dist[x][y]-1] = 'D'
					}
					x = nx
					y = ny
					break
				}
			}
		}
		fmt.Println(string(ans))
	}

}