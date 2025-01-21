package main

import (
	"fmt"
)

func binexp(a, b int) int {
	if b == 0 {
		return 1
	}
	if b%2 == 0 {
		return binexp(a*a, b/2)
	}
	return a * binexp(a*a, b/2)
}

func 
