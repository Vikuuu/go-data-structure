// Implementation of Radix Trie in Go
package datstr

import (
	"fmt"
	"strings"
)

// The Edge
type rdxE struct {
	label string
	next  *rdxN
}

func newRdxE(label string) *rdxE {
	return &rdxE{label: label, next: newRadixNode(false)}
}

func newRdxEWithNext(label string, next *rdxN) *rdxE {
	return &rdxE{label: label, next: next}
}

// Radix Node
type rdxN struct {
	edges  map[rune]*rdxE
	isWord bool
}

func newRadixNode(isWord bool) *rdxN {
	return &rdxN{
		edges:  make(map[rune]*rdxE),
		isWord: isWord,
	}
}

func (r *rdxN) addEdge(label string) {
	r.edges[rune(label[0])] = newRdxE(label)
}

func (r *rdxN) addEdgeWithNext(label string, next *rdxN) {
	r.edges[rune(label[0])] = newRdxEWithNext(label, next)
}

func (r *rdxN) getTransition(transitionChar rune) *rdxE {
	return r.edges[transitionChar]
}

// The Radix Tree
type RadixTrie struct {
	root *rdxN
}

func NewRadixTrie() *RadixTrie {
	return &RadixTrie{root: newRadixNode(false)}
}

func (r *RadixTrie) Search(word string) bool {
	return r.search(r.root, word)
}

func (r *RadixTrie) Insert(word string) error {
	if word == "" {
		return ErrEmptyWord
	}
	r.insert(r.root, word)
	return nil
}

func (r *RadixTrie) Delete(word string) {
	r.delete(r.root, word)
}

func (r *RadixTrie) search(cur *rdxN, word string) bool {
	// base case meow...
	if word == "" {
		return cur.isWord
	}

	transitionChar := getRuneAt(word, 0)
	childEdge := cur.getTransition(transitionChar)

	if childEdge == nil {
		return false
	}

	if !strings.HasPrefix(word, childEdge.label) {
		return false
	}

	// if word == childEdge.label && childEdge.next.isWord {
	// 	return true
	// }

	after, found := strings.CutPrefix(word, childEdge.label)
	if !found {
		return false
	}

	return r.search(childEdge.next, after)
}

func getRuneAt(word string, idx int) rune {
	return rune(word[idx])
}

func (r *RadixTrie) insert(curr *rdxN, word string) {
	if word == "" {
		return
	}

	transitionChar := getRuneAt(word, 0)
	childEdge := curr.getTransition(transitionChar)

	// Case 1: word is not present in the trie
	if childEdge == nil {
		curr.addEdge(word)
		childEdge = curr.getTransition(transitionChar)
		childEdge.next.isWord = true

		return
	}

	// Case 5: the word is already present as a label
	if word == childEdge.label {
		childEdge.next.isWord = true
		return
	}

	cmnPreIdx := r.commonPrefix(word, childEdge.label)

	// Case 2: adding word "slower", "slow" is present
	if cmnPreIdx+1 == len(childEdge.label) {
		suffix := word[cmnPreIdx+1:]
		r.insert(childEdge.next, suffix)
		return
	}

	// Case 3: adding word "slow", "slower" is present
	if cmnPreIdx+1 == len(word) {
		prefix := word[:cmnPreIdx+1]
		suffix := childEdge.label[cmnPreIdx+1:]
		delete(curr.edges, transitionChar)
		// add the prefix word
		curr.addEdge(prefix)
		newChildEdge := curr.getTransition(transitionChar)
		newChildEdge.next.isWord = true
		// recursive call
		r.insert(newChildEdge.next, suffix)
		return
	}

	// Case 4: adding word "waste", "water" is present
	// so the common Prefix idx len will not be equal to
	// the label or the word itself
	prefix := word[:cmnPreIdx+1]
	wordSuffix := word[cmnPreIdx+1:]
	labelSuffix := childEdge.label[cmnPreIdx+1:]
	delete(curr.edges, transitionChar)
	// add the prefix word
	curr.addEdge(prefix)
	newChildEdge := curr.getTransition(transitionChar)
	// if the prefix itself is a complete word
	if len(wordSuffix) == 0 {
		newChildEdge.next.isWord = true
	}
	// reatach the old label suffix under the new prefix
	newChildEdge.next.addEdgeWithNext(labelSuffix, childEdge.next)
	// add the new word's suffix
	if len(wordSuffix) > 0 {
		newChildEdge.next.addEdge(wordSuffix)
		newChildEdge.next.getTransition(getRuneAt(wordSuffix, 0)).next.isWord = true
	}
}

func (r *RadixTrie) delete(cur *rdxN, word string) *rdxN {
	if word == "" {
		return nil
	}

	transitionChar := getRuneAt(word, 0)
	edge := cur.getTransition(transitionChar)

	if edge == nil {
		return nil
	}

	if edge.label == word {
		edge.next.isWord = false
		return cur
	}

	cmnPrefix := r.commonPrefix(word, edge.label)
	// prefix := word[:cmnPrefix+1]
	suffix := word[cmnPrefix+1:]

	del := r.delete(edge.next, suffix)
	if del != nil {
		delete(cur.edges, getRuneAt(suffix, 0))
	}

	return cur
}

func (r *RadixTrie) commonPrefix(s1, s2 string) int {
	if len(s1) < 1 || len(s2) < 1 {
		return -1
	}
	minLen := min(len(s1), len(s2))
	cmnPreLen := -1

	for i := range minLen {
		if s1[i] != s2[i] {
			break
		}
		cmnPreLen++
	}

	return cmnPreLen
}

func (r *RadixTrie) PrintAllWords() {
	r.printAllWords(r.root, "")
}

func (r *RadixTrie) printAllWords(curr *rdxN, result string) {
	if curr.isWord {
		fmt.Println(result)
	}
	for _, v := range curr.edges {
		r.printAllWords(v.next, result+v.label)
	}
}
