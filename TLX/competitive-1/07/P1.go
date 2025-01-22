package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	var n int
	fmt.Scan(&n)
	dp := make([]int, 1001)
	for i := 0; i <= 11; i++ {
		dp[i] = i
	}
	dp[12] = 13
	for i := 13; i <= n; i++ {
		dp[i] = max(dp[i/2] + dp[i/3] + dp[i/4], i) 
	}
	fmt.Print(dp[n])
}
