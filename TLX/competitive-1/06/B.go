package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &a[i])
	}

	fmt.Fscanln(reader)

	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &b[i])
	}

	sort.Ints(a)  
	sort.Ints(b) 
	summ := 0
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d %d\n", a[i], b[i])
		summ += a[i] * b[n-i-1] 
	}
	fmt.Fprintf(writer, "%d", summ)
}
