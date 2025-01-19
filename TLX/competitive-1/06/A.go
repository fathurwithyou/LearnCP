package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, h int
	fmt.Fscanf(reader, "%d %d\n", &n, &h)

	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &b[i])
	}

	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})

	sum := 0
	for i := 0; i < n; i++ {
		sum += b[i]
		if sum >= h {
			fmt.Fprint(writer, i+1)
			return
		}
	}
}
