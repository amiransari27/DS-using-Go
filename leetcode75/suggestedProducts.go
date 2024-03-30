package leetcode75

type TrieNode struct {
	wordEnd  bool
	children [26]*TrieNode
}

func NewTrie() *TrieNode {
	return &TrieNode{}
}
func (t *TrieNode) insertTrieNode(word string) {
	crawl := t
	for _, char := range word {
		idx := char - 'a'
		if crawl.children[idx] == nil {
			crawl.children[idx] = &TrieNode{}
		}

		crawl = crawl.children[idx]
	}
	crawl.wordEnd = true
}

func (t *TrieNode) search(prefix string) []string {
	crawl := t
	res := make([]string, 0)
	for _, char := range prefix {
		idx := char - 'a'
		if crawl.children[idx] == nil {
			return []string{}
		}
		crawl = crawl.children[idx]
	}
	trieDfs(crawl, prefix, &res)
	return res
}

func trieDfs(crawl *TrieNode, prefix string, res *[]string) {
	if len(*res) == 3 {
		return
	}
	if crawl.wordEnd {
		*res = append(*res, prefix)
	}

	for i := 0; i < 26; i++ {
		if crawl.children[i] != nil {
			trieDfs(crawl.children[i], prefix+string(i+'a'), res)
		}
	}
}

func SuggestedProducts(products []string, searchWord string) [][]string {

	trie := NewTrie()
	for _, s := range products {
		trie.insertTrieNode(s)
	}

	result := make([][]string, 0)
	prefix := ""

	for _, char := range searchWord {
		prefix += string(char)
		result = append(result, trie.search(prefix))
	}

	return result
}
