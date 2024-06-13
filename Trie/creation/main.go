package main

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
	// insertedStrings []string
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
	// t.insertedStrings = append(t.insertedStrings, word)
}

func main() {
	trie := newTrie()
	words := []string{"hello", "world", "hey", "hi", "bye"}

	for _, word := range words {
		trie.insert(word)
	}

	// fmt.Println("Inserted strings:", trie.insertedStrings)
}
