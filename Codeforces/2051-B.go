package main

import (
	"fmt"
)

func solve(){
	var n,a,b,c int
	fmt.Scan(&n, &a, &b, &c)
	s := a+b+c
	rem := n%s
	div := n/s
	cnt := 0
	if rem > 0 {
		cnt++
		rem -= a
	}
	if rem > 0 {
		cnt++
		rem -= b
	}
	if rem > 0 {
		cnt++
	}
	fmt.Println(cnt+div*3)
}

func main(){
	var t int
	fmt.Scan(&t)
	for t > 0 {
		t--
		solve()
	}
}
