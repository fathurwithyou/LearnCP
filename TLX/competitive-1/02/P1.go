package main

import (
	"fmt"
)

type Pair struct {
	x, y int
}

func isTriangle(a, b, c Pair) bool {
	if a.x*(b.y-c.y)+b.x*(c.y-a.y)+c.x*(a.y-b.y) == 0 {
		return false
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i].x, &arr[i].y)
	}
	count := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if isTriangle(arr[i], arr[j], arr[k]) {
					count++
				}
			}
		}		
	}
	fmt.Println(count)
}