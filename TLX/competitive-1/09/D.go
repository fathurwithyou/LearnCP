package main

import (
	"fmt"
)
/*
3
3 2 1
3


*/


func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	var n, k int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&k)

	Queue := make([][]int, 0)
	Queue = append(Queue, []int{})

	for i := 0; i < n-k; i++ {
		for 
	}	
}