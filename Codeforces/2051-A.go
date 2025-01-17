package main

import (
	"fmt"
)

func solve(){
	var n int
	fmt.Scan(&n)
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	ans := a[n-1]
	for i := 0; i < n-1; i++ {
		if a[i] > b[i+1] {
			ans += a[i] - b[i+1]
		}
	}
	fmt.Println(ans)
}

func main(){
	var t int
	fmt.Scan(&t)
	for t > 0 {
		t--
		solve()
	}
}
