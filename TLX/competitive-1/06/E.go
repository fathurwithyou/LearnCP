package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Item struct {
	s, d int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	arr := make([]Item, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i].s, &arr[i].d)
		arr[i].d += arr[i].s
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].d < arr[j].d || (arr[i].d == arr[j].d && arr[i].s < arr[j].s)
	})

	end := 0
	ans := 0
	for i := 0; i < n; i++ {
		if arr[i].s >= end {
			ans++
			end = arr[i].d
		}
	}
	fmt.Fprint(writer, ans)


}

/*
	4
	2 5 -> 2 7
	9 7 -> 9 16
	15 6 -> 15 21
	9 3 -> 9 12

	sort by end time
	2 7
	9 12
	15 21
	9 16
*/