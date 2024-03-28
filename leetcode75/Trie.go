package leetcode75

type Trie struct {
	wordEnd  bool
	children [26]*Trie
}

func ConstructorTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	crawl := this

	for _, char := range word {
		idx := char - 'a'
		if crawl.children[idx] == nil {
			crawl.children[idx] = &Trie{}
		}

		crawl = crawl.children[idx]
	}
	crawl.wordEnd = true
}

func (this *Trie) Search(word string) bool {
	crawl := this
	for _, char := range word {
		idx := char - 'a'
		if crawl.children[idx] == nil {
			return false
		}
		crawl = crawl.children[idx]
	}
	return crawl.wordEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	crawl := this

	for _, char := range prefix {
		idx := char - 'a'
		if crawl.children[idx] == nil {
			return false
		}
		crawl = crawl.children[idx]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
