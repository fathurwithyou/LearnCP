package main

import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	for t > 0 {
		t--
		ans := 0
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++{
			fmt.Scan(&arr[i])
		}
		i := 0
		for i < n{
			if arr[i] != 0{
				ans++
				break
			}
			i++
		}
		for i < n{
			if arr[i] == 0{
				j := i
				for j < n && arr[j] == 0{
					j++
				}
				if j != n{
					ans++
				}
				break
			}
			i++
		}
		fmt.Println(ans)
	}
}