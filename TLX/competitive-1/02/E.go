package main

import (
	"bufio"
	"fmt"
	"os"
)

// func gcd(a, b int64) int64 {
// 	if b == 0 {
// 		return a
// 	}
// 	return gcd(b, a%b)
// }

const MOD = 1000000007

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n,a,b int
	fmt.Fscanf(reader, "%d %d %d\n", &n, &a, &b)

	ans := 0
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, b+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 0; i <= b; i++ {
		dp[i][i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= i-1 && j <= b; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			dp[i][j] %= MOD
		}
	}
	for i := a; i <= b; i++ {
		ans += dp[n][i]
		ans %= MOD
	}
	fmt.Fprintln(writer, ans)
}
