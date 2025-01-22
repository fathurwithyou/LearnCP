package main

import (
	"bufio"
	"fmt"
	"os"
)

func UpperBound(arr []int, x int) int {
	l, r := 0, len(arr)
	for l < r {
		m := (l + r) / 2
		if arr[m] <= x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
		a[i] += a[i-1]
	}

	var q, x int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &x)
		idx := UpperBound(a, x)
		if a[idx-1] == x {
			idx--
		}
		fmt.Fprintln(writer, idx)
	}
}