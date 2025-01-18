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

	v := make([]struct {
		length int
		word   string
	}, n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanf(reader, "%s\n", &s)
		v[i] = struct {
			length int
			word   string
		}{length: len(s), word: s}
	}

	sort.Slice(v, func(i, j int) bool {
		return v[i].length < v[j].length || (v[i].length == v[j].length && v[i].word < v[j].word)
	})

	for _, entry := range v {
		fmt.Fprintln(writer, entry.word)
	}
}
