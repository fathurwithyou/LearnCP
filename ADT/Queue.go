package main

import (
	"fmt"
)

type Item struct {
	val interface{}
}

type Queue []Item

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(val interface{}) {
	*q = append(*q, Item{val: val})
}	

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func (q *Queue) Front() interface{} {
	if len(*q) == 0 {
		return nil
	}
	return (*q)[0].val
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Back() interface{} {
	if len(*q) == 0 {
		return nil
	}
	return (*q)[len(*q)-1].val
}

func (q *Queue) Pop() {
	if !(*q).Empty() {
		*q = (*q)[1:]
	}
}

