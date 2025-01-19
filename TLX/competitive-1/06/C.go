package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, d int
	fmt.Fscanf(reader, "%d %d\n", &n, &d)

	h := make([]struct {
		x, y int
	}, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &h[i].x, &h[i].y)
	}

	sort.Slice(h, func(i, j int) bool {
		return h[i].x < h[j].x 
	})

	ans := 0

	for i := 0; i < n; i++ {
		if d >= h[i].x {
			div := d / h[i].x
			ans += min(div, h[i].y)
			d -= min(div, h[i].y) * h[i].x
		} else {
			break
		}
	}

	fmt.Fprint(writer, ans)
}