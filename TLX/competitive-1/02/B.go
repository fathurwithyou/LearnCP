package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Scan(&n)
	tmp := n
	if n == 2 {
		fmt.Println(n)
		return
	}
	if n == 3 {
		fmt.Println(n)
		return
	}
	for i := 2; i <= n && tmp > 1; i++ {
		if tmp % i == 0 {
			fmt.Print(i)
			cnt := 0
			for tmp % i == 0 {
				tmp /= i
				cnt++
			}
			if cnt > 1 {
				fmt.Print("^", cnt)
			}
			if tmp > 1 {
				fmt.Print(" x ")
			}
		}
	}
}