package main

import (
	"fmt"
)

func main() {
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)
	arr := [14]int{}
	arr[a]++
	arr[b]++
	arr[c]++
	arr[d]++
	two := 0
	three := 0
	for i:=0;i<=13;i++{
		if arr[i] == 3{
			three++
		}
		if arr[i] == 2{
			two++
		}
	}
	if (three == 1 || two == 2){
		fmt.Println("Yes")
	} else{
		fmt.Println("No")
	}
}