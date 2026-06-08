package top150i

import "fmt"

type TNode struct {
	children map[rune]*TNode
	eow      bool
}

func (tn *TNode) Insert(key string) {
	if tn.children == nil {
		tn.children = make(map[rune]*TNode)
	}
	curr := tn
	for _, char := range key {
		if curr.children[char] == nil {
			curr.children[char] = &TNode{
				children: make(map[rune]*TNode),
			}
		}
		curr = curr.children[char]
	}
	curr.SetEOW()
}

func (tn *TNode) Search(key string) bool {
	curr := tn
	for _, char := range key {
		if curr.children[char] == nil {
			return false
		}
		curr = curr.children[char]
	}
	return curr.eow
}

func (tn *TNode) SetEOW() {
	tn.eow = true
}

func TestTrieDS() {
	words := []string{"the", "a", "there", "their", "any", "apple"}

	rootTrie := &TNode{}
	for _, w := range words {
		rootTrie.Insert(w)
	}

	fmt.Println(rootTrie.Search("there"))
	fmt.Println(rootTrie.Search("the"))
	fmt.Println(rootTrie.Search("appe"))
}
