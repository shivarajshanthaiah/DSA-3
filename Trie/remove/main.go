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

func (t *Trie) remove(s string) {
	t.deleterecursive(t.root, s, 0)
}

func (t *Trie) deleterecursive(node *TrieNode, s string, i int) *TrieNode {
	if node == nil {
		return nil
	}
	if i == len(s) {
		if node.isEnd {
			node.isEnd = false
		}
		if len(node.children) == 0 {
			return nil
		}
		return node
	}

	char := rune(s[i])
	node.children[char] = t.deleterecursive(node.children[char], s, i+1)
	if node.children[char] == nil {
		delete(node.children, char)
	}

	if len(node.children) == 0 && !node.isEnd {
		return nil
	}
	return node
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

	trie.remove("hello")

	word := "hey"
	fmt.Println(trie.search(word))

}
