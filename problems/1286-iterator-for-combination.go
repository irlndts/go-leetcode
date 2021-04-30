package problems

type CombinationIterator struct {
	List   []string
	length int
}

var list []string

func Constructor(characters string, combinationLength int) CombinationIterator {
	list = []string{}
	for i := 0; i < len(characters); i++ {
		help(characters[i:], string(characters[i]), combinationLength)
	}
	return CombinationIterator{
		List:   list,
		length: len(list),
	}
}

func help(s, cur string, length int) {
	if len(cur) == length {
		list = append(list, cur)
		return
	}

	for i := 1; i < len(s); i++ {
		help(s[i:], cur+string(s[i]), length)
	}
}

func (this *CombinationIterator) Next() string {
	this.length--
	return this.List[len(this.List)-this.length-1]
}

func (this *CombinationIterator) HasNext() bool {
	return this.length != 0
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
