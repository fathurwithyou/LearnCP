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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var q, x, y int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &x, &y)
		lower := UpperBound(a, x)
		upper := UpperBound(a, y)
		fmt.Fprintln(writer, upper-lower)
	}
}