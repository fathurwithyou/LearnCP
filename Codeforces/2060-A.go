package main

import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0;i<t;i++{
		var a,b,c,d int
		fmt.Scan(&a,&b,&c,&d)
		arr := []int{a+b,c-b,d-c}
		maxi := 1
		for j:=0;j<3;j++{
			tmp := 1
			for k:=j+1;k<3;k++{
				if arr[j] == arr[k]{
					tmp++
				}
			}
			if tmp > maxi{
				maxi = tmp
			}
		}
		fmt.Println(maxi)
	}
}