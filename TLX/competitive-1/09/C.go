package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func search(a, b string, par map[string]string) bool {
	for a != "" {
		if a == b {
			return true
		}
		a = par[a]
	}
	return false
}

func print(a, b string, par map[string]string) {
	if a == b {
		fmt.Println(a)
		return
	}
	print(par[a], b, par)
	fmt.Println(a)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)	
	par := make(map[string]string)
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		par[parts[1]] = parts[0]
	}
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	a, b := parts[0], parts[1]
	if search(a, b, par) {
		print(a, b, par)
	} else if search(b, a, par) {
		print(b, a, par)
	} else {
		fmt.Print("TIDAK MUNGKIN")
	}
}