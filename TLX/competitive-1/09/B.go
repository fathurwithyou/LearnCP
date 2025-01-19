package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	moveX = []int{0, 0, 1, -1}
	moveY = []int{1, -1, 0, 0}
	area  int
	seri  bool
)

func findSol(x, y int, arr [][]rune, m, n int) {
	arr[y][x] = '#' // Mark the cell as visited
	for i := 0; i < 4; i++ {
		newX := x + moveX[i]
		newY := y + moveY[i]
		if newX >= 0 && newX < m && newY >= 0 && newY < n && arr[newY][newX] != '#' && !seri {
			if arr[newY][newX] == 'K' {
				seri = true
				return
			}
			area++
			findSol(newX, newY, arr, m, n)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var m, n int
	fmt.Fscanf(reader, "%d %d\n", &m, &n)

	arr := make([][]rune, n)
	var bx, by, kx, ky int = -1, -1, -1, -1

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = line[:len(line)-1] // Remove newline
		arr[i] = []rune(line)

		if bx == -1 && by == -1 {
			for j, ch := range arr[i] {
				if ch == 'B' {
					bx, by = j, i
				}
			}
		}

		if kx == -1 && ky == -1 {
			for j, ch := range arr[i] {
				if ch == 'K' {
					kx, ky = j, i
				}
			}
		}
	}

	findSol(bx, by, arr, m, n)
	areaB := area
	if seri {
		fmt.Fprintln(writer, "SERI")
		return
	}

	area = 0
	findSol(kx, ky, arr, m, n)
	areaK := area

	if areaB > areaK {
		fmt.Fprintf(writer, "B %d\n", areaB-areaK)
	} else if areaB == areaK {
		fmt.Fprintln(writer, "SERI")
	} else {
		fmt.Fprintf(writer, "K %d\n", areaK-areaB)
	}
}
