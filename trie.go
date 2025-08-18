package datstr

import (
	"slices"
	"strings"
)

type trieNode struct {
	value  rune
	isWord bool
	links  []*trieNode
	parent *trieNode
}

func NewTrieNode(val rune) *trieNode {
	return &trieNode{
		value:  val,
		isWord: false,
		links:  make([]*trieNode, 26),
		parent: nil,
	}
}

type Trie struct {
	head *trieNode
}

func NewTrie(length int) *Trie {
	r := ' '
	return &Trie{
		head: NewTrieNode(r),
	}
}

func (t *Trie) Insert(val string, isWord bool) {
	cur := t.head
	for _, v := range strings.ToLower(val) {
		v = v - 'a'
		node := cur.links[v]
		if node != nil {
			continue
		}
		newNode := NewTrieNode(v)
		newNode.parent = cur
		node = newNode
		cur = node
	}
	cur.isWord = isWord
}

func (t *Trie) Search(val string) bool {
	cur := t.head
	for _, v := range strings.ToLower(val) {
		v = v - 'a'
		node := cur.links[v]
		if node == nil {
			return false
		}
		cur = node.links[v]
	}
	return true
}

func (t *Trie) Delete(val string) {
	cur := t.head
	stack := []*trieNode{}
	for _, v := range strings.ToLower(val) {
		v = v - 'a'
		node := cur.links[v]
		stack = append(stack, node)
	}
	slices.Reverse(stack)
	for _, s := range stack {
		// if slice not contains all the nil value
		allNilVal := slices.ContainsFunc(s.links, func(t *trieNode) bool {
			return t != nil
		})
		if allNilVal {
			parent := s.parent
			v := s.value - 'a'
			parent.links[v] = nil
		}
	}
}
