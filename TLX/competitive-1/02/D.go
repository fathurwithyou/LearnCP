package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c, d int64
	fmt.Fscanf(reader, "%d %d\n%d %d\n", &a, &b, &c, &d)

	lcm := (b * d) / gcd(b, d) 
	x := (lcm/b)*a + (lcm/d)*c 
	fpb := gcd(x, lcm)         

	fmt.Fprintf(writer, "%d %d", x/fpb, lcm/fpb)
}
