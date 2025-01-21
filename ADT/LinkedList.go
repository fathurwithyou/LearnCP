package main

import (
	"fmt"
)

type Node struct {
	val   int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func NewNode(val int) *Node {
	return &Node{val: val}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Insert(val int) {
	if ll.head == nil {
		ll.head = NewNode(val)
		return
	}
	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = NewNode(val)
}

func (ll *LinkedList) Delete(val int) {
	if ll.head == nil {
		return
	}
	if ll.head.val == val {
		ll.head = ll.head.next
		return
	}
	cur := ll.head
	for cur.next != nil {
		if cur.next.val == val {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (ll *LinkedList) Search(val int) bool {
	cur := ll.head
	for cur != nil {
		if cur.val == val {
			return true
		}
		cur = cur.next
	}
	return false
}

func (ll *LinkedList) Print() {
	cur := ll.head
	for cur != nil {
		fmt.Print(cur.val, " ")
		cur = cur.next
	}
	fmt.Println()
}

func (ll *LinkedList) Reverse() {
	var prev *Node
	cur := ll.head
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	ll.head = prev
}

func (ll *LinkedList) Remove(val int) {
	if ll.head == nil {
		return
	}
	if ll.head.val == val {
		ll.head = ll.head.next
	}
	cur := ll.head
	for cur.next != nil {
		if cur.next.val == val {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}
}

