// package main

// import (
// 	"fmt"
// )

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// func main() {
// 	var k int
// 	var s, t string
// 	fmt.Scan(&k, &s, &t)
// 	if abs(len(s)-len(t)) > 1 {
// 		fmt.Print("Yes")
// 		return
// 	}
// 	if len(s) == len(t) {
// 		cnt := 0
// 		for i := 0; i < len(s); i++ {
// 			if s[i] != t[i] {
// 				cnt++
// 			}
// 		}
// 		if cnt <= 1{
// 			fmt.Print("Yes")
// 		} else {
// 			fmt.Print("No")
// 		}
// 		return
// 	}
// 	s_idx := 0
// 	t_idx := 0
// 	for i := 0; i < min(len(s), len(t)); i++ {
// 		if s[i] != t[i] {
// 			s_idx = i
// 			t_idx = i
// 			break
// 		}
// 	}
// 	if len(s) < len(t) {
// 		t_idx++
// 	} else {
// 		s_idx++
// 	}
// 	for i := s_idx; i < len(s); i++ {
// 		if s[i] != t[t_idx] {
// 			fmt.Print("No")
// 			return
// 		}
// 		t_idx++
// 	}
// 	fmt.Print("Yes")



// }