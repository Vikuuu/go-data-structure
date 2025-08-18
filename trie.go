package datstr

import (
	"slices"
	"strings"
)

type trieNode struct {
	value  string
	isWord bool
	links  []*trieNode
	parent *trieNode
}

func NewTrieNode(val string) *trieNode {
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
	return &Trie{
		head: NewTrieNode(""),
	}
}

func (t *Trie) Insert(val string) {
	cur := t.head
	for _, v := range strings.ToLower(val) {
		nv := v - 'a'
		childNode := cur.links[nv]
		if childNode == nil {
			childNode = NewTrieNode(string(v))
			cur.links[nv] = childNode
		}
		childNode.parent = cur
		cur = childNode
	}
	cur.isWord = true
}

func (t *Trie) Search(val string) bool {
	cur := t.head
	for _, v := range strings.ToLower(val) {
		nv := v - 'a'
		node := cur.links[nv]
		if node == nil {
			return false
		}
		cur = node
	}
	return cur.isWord
}

func (t *Trie) Delete(val string) {
	cur := t.head
	stack := []*trieNode{}
	for _, v := range strings.ToLower(val) {
		v = v - 'a'
		node := cur.links[v]
		if node == nil { // word does not exists
			return
		}
		stack = append(stack, node)
		cur = node
	}

	if !cur.isWord {
		return
	}
	cur.isWord = false
	slices.Reverse(stack)
	for _, s := range stack {
		var hasChildren bool
		for _, l := range s.links {
			if l != nil {
				hasChildren = true
			}
		}
		if s.isWord || hasChildren {
			break
		}
		parent := s.parent
		if parent != nil {
			v := rune(s.value[0]) - 'a'
			parent.links[v] = nil
		}
	}
}
