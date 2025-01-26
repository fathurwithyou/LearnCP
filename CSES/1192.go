package main

import (
	"fmt"
) 

var vis [1001][1001]bool
var arr [1001]string

var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}
var n, m int

func dfs(i, j int){
	vis[i][j] = true
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if x >= 0 && x < n && y >= 0 && y < m && arr[x][y] == '.' && !vis[x][y] {
			dfs(x, y)
		}
	}
}

func main(){
	fmt.Scan(&n, &m)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	cnt := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == '.' && !vis[i][j] {
				cnt++
				dfs(i, j)			
			}
		}
	}

	fmt.Print(cnt)
}