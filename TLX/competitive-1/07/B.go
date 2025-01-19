package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscanf(reader, "%d %d", &n, &k)

	w := make([]int, k+1)
	h := make([]int, k+1)
	w[0] = 0
	h[0] = 0

	for i := 1; i <= k; i++ {
		fmt.Fscanf(reader, "\n%d %d", &w[i], &h[i])
	}

	dp := make([][]int, k+1)
	backtrack := make([][]pair, k+1)

	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
		backtrack[i] = make([]pair, n+1)
	}		

	for i := 1; i <= k; i++ {
		for j := 0; j <= n; j++ {
			if j >= w[i] && dp[i-1][j] < dp[i-1][j-w[i]]+h[i] {
				dp[i][j] = dp[i-1][j-w[i]] + h[i]
				backtrack[i][j] = pair{i - 1, j - w[i]}
			} else {
				dp[i][j] = dp[i-1][j]
				backtrack[i][j] = pair{i - 1, j}
			}
		}
	}

	x := n
	for i := 1; i <= n; i++ {
		if dp[k][i] == dp[k][n] {
			x = i
			break
		}
	}

	ans := []int{}
	for i := k; i >= 1; i-- {
		if backtrack[i][x].second == x {
			continue
		} else {
			ans = append(ans, i)
			x = backtrack[i][x].second
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprintln(writer, ans[i])
	}
}

type pair struct {
	first, second int
}
