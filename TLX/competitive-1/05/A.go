package main

import (
	"fmt"
)

const mod = 1e6
var big bool

func binexp(a, b int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = res * a
			if res >= mod {
				big = true
				res %= mod
			}
		}
		a = a * a
		a %= mod
		b /= 2
	}
	return res
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res := binexp(a, b)
	if res == 0 {
		fmt.Println("000000")
	} else if big {
		fmt.Printf("%06d\n", res)
	} else {
		fmt.Println(res)
	}
}