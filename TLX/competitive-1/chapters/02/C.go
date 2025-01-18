package main

import (
	"bufio"
	"fmt"
	"os"
)

var arr [100005]int
var idx int

func sieve(n int) {
	elim := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !elim[i] {
			arr[idx] = i
			idx++
			j := i * i
			for j <= n {
				elim[j] = true
				j += i
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	sieve(1000000)

	for i := 0; i < t; i++ {
		var k int
		fmt.Fscanf(reader, "%d\n", &k)
		fmt.Fprintln(writer, arr[k-1])
	}
}
