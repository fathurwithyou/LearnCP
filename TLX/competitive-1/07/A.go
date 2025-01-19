package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solve() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscanf(reader, "%d\n", &n)

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d", &arr[i])
	}

	fmt.Fscanf(reader, "%d\n", &k)

	dp := make([]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = math.MaxInt32 // Initialize to a large value
	}

	dp[0] = 0
	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {
			if i-arr[j] >= 0 {
				dp[i] = min(dp[i], dp[i-arr[j]]+1)
			}
		}
	}

	if dp[k] == math.MaxInt32 {
		fmt.Fprintln(writer, -1) // If no solution is found
	} else {
		fmt.Fprintln(writer, dp[k])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	solve()
}
