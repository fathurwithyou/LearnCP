package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		t--
		var n, m int
		fmt.Scan(&n, &m)
		fmt.Println(max(n, m) + 1)
	}
}
