package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Item struct {
	c byte
	v int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	s, _ := reader.ReadString('\n')
	s = s[:n] // Trim the newline character

	tab := make([]Item, 26)

	for i := 0; i < n; i++ {
		tab[s[i]-'A'].c = s[i]
		tab[s[i]-'A'].v++
	}

	sort.Slice(tab, func(i, j int) bool {
		return tab[i].v > tab[j].v
	})

	if tab[2].v == 0 {
		fmt.Print(-1)
	} else {
		cnt := tab[2].v * 3
		if tab[1].v > tab[2].v {
			cnt++
		}
		if tab[0].v > tab[2].v {
			cnt++
		}
		fmt.Print(cnt)
	}
}