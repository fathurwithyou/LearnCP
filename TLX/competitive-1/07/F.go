package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		if i % 4 == 0  && i != 4{
			dp[i]++
		} 
	}
	fmt.Print(dp[n])

}
/*
1 
2 -> 1

1 3 
2 4 -> 2

1 3 5
2 4 6 -> 3

1 2 3 4
8 7 6 5



dp[i] = dp[i-1] + dp[i-2]
*/