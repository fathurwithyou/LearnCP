package main

import (
	"fmt"
)

type Stack []int

func (s Stack) Len() int {
	return len(s)
}
func (s *Stack) Push(x int) {
	*s = append(*s, x)
}
func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}
func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}
func (s Stack) Empty() bool {
	return s.Len() == 0
}

