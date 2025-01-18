package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solve() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, x int
	fmt.Fscanf(reader, "%d %d\n", &n, &x)

	res := make([]int, n)
	mini := 200000
	ans := 200000

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &res[i])
		mini = int(math.Min(float64(mini), math.Abs(float64(res[i]-x))))
	}

	for i := 0; i < n; i++ {
		if int(math.Abs(float64(x-res[i]))) == mini {
			ans = int(math.Min(float64(ans), float64(res[i])))
		}
	}

	fmt.Fprintf(writer, "%d\n", ans)
}

func main() {
	solve()
}