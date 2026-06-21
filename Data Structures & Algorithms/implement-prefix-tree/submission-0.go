
type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}
type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	return PrefixTree{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (this *PrefixTree) Insert(word string) {
	current := this.root
	for _, char := range word {
		if current.children[char] == nil {
			current.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		current = current.children[char]
	}
	current.endOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	current := this.root
	for _, char := range word {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return current.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	current := this.root
	for _, char := range prefix {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return true
}
