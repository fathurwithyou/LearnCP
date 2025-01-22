package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	w, c, v float64
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0])
	x, _ := strconv.ParseFloat(parts[1], 64)

	arr := make([]Item, n)
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	weights := strings.Split(line, " ")
	for i := 0; i < n; i++ {
		arr[i].w, _ = strconv.ParseFloat(weights[i], 64)
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	costs := strings.Split(line, " ")
	for i := 0; i < n; i++ {
		arr[i].c, _ = strconv.ParseFloat(costs[i], 64)
		arr[i].v = arr[i].c / arr[i].w
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].v > arr[j].v
	})

	var total float64
	for i := 0; i < n; i++ {
		if x > 0 {
			total += min(x, arr[i].w) * arr[i].v
			x -= min(x, arr[i].w)
		} else {
			break
		}
	}

	fmt.Printf("%.5f\n", total)
}