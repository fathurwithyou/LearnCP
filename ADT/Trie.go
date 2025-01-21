package main

import (
	"fmt"
)

ALPHABET_SIZE := 26

type Node struct {
	val      int
	children [ALPHABET_SIZE]*Node
}

type Trie struct {
	root *Node
}

func NewNode(val int) *Node {
	return &Node{val: val}
}

func NewTrie() *Trie {
	return &Trie{root: NewNode(0)}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			cur.children[c-'a'] = NewNode(0)
		}
		cur = cur.children[c-'a']
	}
	cur.val = 1
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			return false
		}
		cur = cur.children[c-'a']
	}
	return cur.val == 1
}