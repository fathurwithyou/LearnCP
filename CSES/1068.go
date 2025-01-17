package main

import (
	"fmt"
)

func main(){
	var n int32
	fmt.Scan(&n)
	for n != 1 {
		fmt.Print(n, " ")
		if n & 1 == 1{
			n = n*3+1
		} else{
			n /= 2	
		}
	}
	fmt.Print(1)
}