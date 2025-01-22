package main

import (
	"fmt"
)

type Tuple struct {
	h, k, d int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	arr := make([]Tuple, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i].h, &arr[i].k, &arr[i].d)
	}
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, m+1)
	}
	/*
	i for item
	j for money
	*/
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			dp[1][j] = dp[0][j]
			if j >= arr[i].h {
				if arr[i].d == 1 {
					dp[1][j] = max(dp[1][j], dp[0][j-arr[i].h]+arr[i].k)
				} else {
					dp[1][j] = max(dp[1][j], dp[0][j-arr[i].h]+arr[i].k)
					dp[1][j] = max(dp[1][j], dp[1][j-arr[i].h]+arr[i].k)
				}
			}
			// fmt.Print(dp[1][j], " ")
			// dp[0][j] = dp[1][j]
		}
		for j := 0; j <= m; j++ {
			dp[0][j] = dp[1][j]
		}
		// fmt.Println()
	}
	fmt.Print(dp[1][m])
}