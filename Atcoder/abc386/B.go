package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	cnt := 0
	for i:=0;i<len(s);i++{
		if s[i] == '0'{
			cnt++
			if i+1 < len(s) && s[i+1] == '0'{
				i++
			}
		} else{
			cnt++
		}
	}
	fmt.Print(cnt)
}