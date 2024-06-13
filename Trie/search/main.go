package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) insert(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func main() {
	trie := newTrie()
	words := []string{"hello", "world", "hey", "hi", "bye"}

	for _, word := range words {
		trie.insert(word)
	}

	// searchWords := []string{"hello", "apple", "hey"}
	// for _, word := range searchWords {
	// 	if trie.search(word) {
	// 		fmt.Printf("%s found\n", word)
	// 	} else {
	// 		fmt.Printf("%s not found\n", word)
	// 	}
	// }
	word := "hello"
	fmt.Println(trie.search(word))
}
