package problems

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

type WordDictionary struct {
	Head *Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		Head: NewNode(""),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
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

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
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

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
