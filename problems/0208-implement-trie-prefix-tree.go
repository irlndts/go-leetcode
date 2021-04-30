package problems

type Trie struct {
	Head *Node
}

type Node struct {
	Val   string
	End   bool
	Nodes map[string]*Node
}

func NewNode(val string) *Node {
	return &Node{
		Val:   val,
		End:   false,
		Nodes: map[string]*Node{},
	}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Head: NewNode(""),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	n := this.Head
	for _, r := range word {
		s := string(r)
		if _, ok := n.Nodes[s]; !ok {
			n.Nodes[s] = NewNode(s)
		}
		n = n.Nodes[s]
	}
	n.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	n := this.Head
	return search(word, n)
}

func search(word string, n *Node) bool {
	if word == "" {
		if !n.End {
			return false
		}
		return true
	}

	s := string(word[0])

	if s == "." {
		for k := range n.Nodes {
			if search(word[1:], n.Nodes[k]) {
				return true
			}
		}
		return false
	}

	if _, ok := n.Nodes[s]; !ok {
		return false
	}
	return search(word[1:], n.Nodes[s])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	n := this.Head
	for _, r := range prefix {
		s := string(r)
		if _, ok := n.Nodes[s]; !ok {
			return false
		}
		n = n.Nodes[s]
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
