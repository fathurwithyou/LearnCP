package main

import (
	"fmt"
)

func main(){
	var b, k int
	fmt.Scan(&b, &k)
	k -= b
	for i := 60; i >= 0; i--{
		if k & (1<<i) != 0{
			fmt.Println(1<<i)
		}
	}
}