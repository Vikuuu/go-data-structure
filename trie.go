package datstr

import (
	"errors"
	"strings"
)

var (
	ErrNoSuchWord  = errors.New("no such word in trie")
	ErrNotWord     = errors.New("not word")
	ErrEmptyWord   = errors.New("empty word")
	ErrInvalidRune = errors.New("invalid character")
)

type trieNode struct {
	value    rune
	children [26]*trieNode
	parent   *trieNode
	isWord   bool
}

func NewTrieNode(val rune) *trieNode {
	return &trieNode{
		value: val,
	}
}

type Trie struct {
	head *trieNode
}

func NewTrie() *Trie {
	var r rune
	return &Trie{head: NewTrieNode(r)}
}

func getIndex(r rune) (int, error) {
	if r < 'a' || r > 'z' {
		return -1, ErrInvalidRune
	}
	return int(r - 'a'), nil
}

func (t *Trie) Insert(word string) error {
	if word == "" {
		return ErrEmptyWord
	}
	cur := t.head
	for _, r := range strings.ToLower(word) {
		idx, err := getIndex(r)
		if err != nil {
			return err
		}

		child := cur.children[idx]
		if child == nil {
			child = NewTrieNode(r)
			child.parent = cur
			cur.children[idx] = child
		}
		cur = child
	}
	cur.isWord = true
	return nil
}

func (t *Trie) Search(word string) error {
	if word == "" {
		return ErrEmptyWord
	}
	cur := t.head
	for _, r := range strings.ToLower(word) {
		idx, err := getIndex(r)
		if err != nil {
			return err
		}

		child := cur.children[idx]
		if child == nil {
			return ErrNoSuchWord
		}
		cur = child
	}
	if cur.isWord {
		return nil
	}
	return ErrNoSuchWord
}

func (t *Trie) Delete(word string) error {
	if word == "" {
		return ErrNotWord
	}

	cur := t.head
	for _, r := range strings.ToLower(word) {
		idx, err := getIndex(r)
		if err != nil {
			return err
		}

		child := cur.children[idx]
		if child == nil {
			return ErrNoSuchWord
		}
		cur = child
	}

	if !cur.isWord {
		return ErrNotWord
	}
	cur.isWord = false

	parent := cur.parent
	for parent != nil {
		if t.hasChildren(parent) || cur.isWord {
			break
		}
		idx, err := getIndex(cur.value)
		if err != nil {
			return err
		}

		parent.children[idx] = nil
		cur = parent
		parent = parent.parent
	}
	return nil
}

func (t *Trie) hasChildren(parent *trieNode) bool {
	count := 0
	for _, pointer := range parent.children {
		if pointer != nil {
			count++
		}
	}
	if count > 1 {
		return true
	}
	return false
}
