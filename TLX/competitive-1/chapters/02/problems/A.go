package main

import (
	"fmt"
)

func check(n int) bool{
	cnt := 0
	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return true
}

func main(){
	var n int
	fmt.Scan(&n)
	for n > 0 {
		n--
		var x int
		fmt.Scan(&x)
		if check(x) {
			fmt.Println("YA")
		} else {
			fmt.Println("BUKAN")
		}
	}
}