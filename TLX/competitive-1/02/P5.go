package main

import (
	"fmt"
	"math"
)

func main(){
	var b,c,d int
	cnt := 0
	fmt.Scan(&b,&c,&d)
	if c == d && b > c{
		fmt.Println(b-c)
		return
	}
	for i := 1; i*i <= c-d; i++ {
		
	}
	fmt.Println(cnt)
}	

/*
18 13 3
10 % i == 0
13 % 10 == 3
13 % 15 == 3
13 % 20 == 3


for i <- 1 to b do
	if c % i == d then
		cnt++
	end if
end for

c%i == d === (c-d)%i == 0
*/