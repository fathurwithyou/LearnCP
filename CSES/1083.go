package main 

import (
	"fmt"
)

func main(){
	var n int32
	fmt.Scan(&n)
	arr := make([]bool, n)
	for i := 0; i < n-1; i++ {
		var a int
		fmt.Scan(&a)
		arr[a-1] = true
	}
	for i := 0; i < n; i++ {
		if !arr[i] {
			fmt.Print(i+1)
			break
		}
	}
}