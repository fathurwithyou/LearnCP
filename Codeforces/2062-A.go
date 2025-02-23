package main

import (
	"fmt"
	// "math"
)

func main(){
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var s string
		fmt.Scan(&s)
		cnt := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '1'{
				cnt++
			}
		}
		fmt.Println(cnt)
		t--
	}
}