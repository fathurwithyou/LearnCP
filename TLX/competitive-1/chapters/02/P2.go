package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int32) int32 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int32
	fmt.Fscanf(reader, "%d\n", &n)

	var ans int32 = 1
	for i := int32(0); i < n; i++ {
		var temp int32
		fmt.Fscanf(reader, "%d\n", &temp)
		ans = ans * temp / gcd(ans, temp)
	}

	fmt.Fprintf(writer, "%d\n", ans)
}
