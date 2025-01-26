package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	var x int
	fmt.Scan(&n, &x)
	fmt.Println(big.NewInt(n).Text(x))
}