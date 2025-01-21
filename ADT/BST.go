package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{val: val}
}

func insertNode(root *Node, val int) *Node {
	if root == nil {
		return NewNode(val)
	}
	if val < root.val {
		root.left = insertNode(root.left, val)
	} else {
		root.right = insertNode(root.right, val)
	}
	return root
}

func searchNode(root *Node, val int) bool {
	if root == nil {
		return false
	}
	if root.val == val {
		return true
	}
	if val < root.val {
		return searchNode(root.left, val)
	}
	return searchNode(root.right, val)
}

func buildBSTFromSortedArray(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	root := NewNode(arr[mid])
	root.left = buildBSTFromSortedArray(arr[:mid])
	root.right = buildBSTFromSortedArray(arr[mid+1:])
	return root
}

func printHelperTree(root *Node, tab int, now int){
	if root == nil {
		return
	}
	for i := 0; i < now; i++ {
		fmt.Print(" ")
	}
	fmt.Println(root.val)
	printHelperTree(root.left, tab, now+tab)
	printHelperTree(root.right, tab, now+tab)
}

func printTree(root *Node, tab int) {
	printHelperTree(root, tab, 0)
}
