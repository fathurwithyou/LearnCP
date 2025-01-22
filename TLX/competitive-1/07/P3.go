package main

import (
	"fmt"
)

type Tuple struct {
	x1, y1, x2, y2 int
}

func main() {
	var v, h, n int
	fmt.Scan(&v, &h, &n)
	dp := make([][]int, v+1)
	for i := 0; i <= v; i++ {
		dp[i] = make([]int, h+4)
	}
	arr := make([]Tuple, n)
	for i := 0; i < n; i++ {
		var x, y, a int
		fmt.Scan(&a, &x, &b, &y)
		arr[i] = Tuple{a, x, b, y}		
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].x1 < arr[j].x1 || (arr[i].x1 == arr[j].x1 && arr[i].y1 < arr[j].y1)
	})

	for i := 0; i < n; i++{
		
	}

	ans := 0
	for i := 1; i <= h; i++ {
		ans = max(ans, dp[v][i])
	}
	fmt.Print(ans)
}
/*
x x
x x

 x




*/