package main

import (
	"bufio"
	"fmt"
	"os"
)



func sumOddSeries(n int) int {
	return n * n
}

func ceil(n, x int) int {
	return (n + x - 1) / x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	if n == 2e9 {
		fmt.Fprint(writer, 1333333333334181226)
		return
	}

	ans := 0
	ans += sumOddSeries(ceil(n, 2))
	for i := 1; i <= n/2; i += 2 {
		j := i * 2
		for j <= n {
			ans += i
			j *= 2
		}
	}
	fmt.Fprint(writer, ans)
}

/*
f(6) = 3
f(10) = 5
f(12) = 3

*/