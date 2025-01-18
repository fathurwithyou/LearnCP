package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%f\n", &arr[i])
	}

	sort.Float64s(arr) // Sort the array

	if n%2 == 0 {
		median := (arr[n/2] + arr[n/2-1]) / 2.0
		fmt.Fprintf(writer, "%.1f\n", median)
	} else {
		median := arr[n/2]
		fmt.Fprintf(writer, "%.1f\n", median)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		solve()
	}
}
