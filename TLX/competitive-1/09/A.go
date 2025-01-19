package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 100005

var (
	adj  = make([][]int, MAXN)
	vis  = make([]bool, MAXN)
)

func dfs(u int) {
	vis[u] = true
	for _, v := range adj[u] {
		if !vis[v] {
			dfs(v)
		}
	}
}

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	// Reset adjacency list and visited array for each test case
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
		vis[i] = false
	}

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscanf(reader, "%d %d\n", &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if !vis[i] {
			dfs(i)
			ans++
		}
	}

	fmt.Fprintln(writer, ans)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solve(reader, writer)
}
