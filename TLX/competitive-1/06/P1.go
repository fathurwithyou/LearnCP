package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	nm := strings.Fields(line)
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	an := make([]int, n)
	bn := make([]int, 10001)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &an[i])
	}

	for i := 0; i < m; i++ {
		var b int
		fmt.Fscan(reader, &b)
		bn[b]++
	}

	ans := 0
	sort.Ints(an)
	for i := 0; i < n; i++ {
		if bn[an[i]] > 0 {
			ans++
			bn[an[i]]--
		} else if bn[an[i]+1] > 0 {
			ans++
			bn[an[i]+1]--
		}
	}

	fmt.Fprint(writer, ans)
}