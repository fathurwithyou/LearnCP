package main

import (
	"fmt"
)

var n, m, x int

func vertikal(arr *[101][101]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m/2; j++ {
			(*arr)[i][j], (*arr)[i][m-j-1] = (*arr)[i][m-j-1], (*arr)[i][j]
		}
	}
}

func horizontal(arr *[101][101]int) {
	for i := 0; i < n/2; i++ {
		for j := 0; j < m; j++ {
			(*arr)[i][j], (*arr)[n-i-1][j] = (*arr)[n-i-1][j], (*arr)[i][j]
		}
	}
}

func rotate90(arr *[101][101]int) {
	var temp [101][101]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp[i][j] = (*arr)[n-j-1][i]
		}
	}
	n, m = m, n
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			(*arr)[i][j] = temp[i][j]
		}
	}
}

func rotate180(arr *[101][101]int) {
	vertikal(arr)
	horizontal(arr)
}

func rotate270(arr *[101][101]int) {
	var temp [101][101]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp[i][j] = (*arr)[j][m-i-1]
		}
	}
	n, m = m, n
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			(*arr)[i][j] = temp[i][j]
		}
	}
}

func main() {
	fmt.Scan(&n, &m, &x)

	var mat [101][101]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	for x > 0 {
		x--
		var s string
		fmt.Scan(&s)
		if s == "_" {
			horizontal(&mat)
		} else if s == "|" {
			vertikal(&mat)
		} else if s == "90" {
			rotate90(&mat)
		} else if s == "180" {
			rotate180(&mat)
		} else if s == "270" {
			rotate270(&mat)
		}
		
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}
}
