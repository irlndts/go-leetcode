package problems

type StreamChecker struct {
	Request []byte
	Items   map[byte]*Item
}

type Item struct {
	Val   byte
	Items map[byte]*Item

	End bool
}

func Constructor(words []string) StreamChecker {
	checker := StreamChecker{Items: map[byte]*Item{}, Request: []byte{}}
	for _, w := range words {
		start := w[len(w)-1]
		item, ok := checker.Items[start]
		if !ok {
			item = &Item{Val: byte(start), Items: map[byte]*Item{}}
			checker.Items[start] = item
		}

		if len(w) == 1 {
			checker.Items[start].End = true
		}

		for i := len(w) - 2; i >= 0; i-- {
			newItem, ok := item.Items[w[i]]
			if !ok {
				newItem = &Item{Val: byte(w[i]), Items: map[byte]*Item{}}
				item.Items[w[i]] = newItem
			}
			item = newItem
		}
		item.End = true
	}

	return checker
}

func (this *StreamChecker) Query(letter byte) bool {
	this.Request = append(this.Request, letter)
	return this.exists()
}

func (this *StreamChecker) exists() bool {
	first, ok := this.Items[this.Request[len(this.Request)-1]]
	if !ok {
		return false
	}

	if first.End {
		return true
	}

	for i := len(this.Request) - 2; i >= 0; i-- {
		v, ok := first.Items[this.Request[i]]
		if !ok {
			return false
		}
		if v.End {
			return true
		}
		first = v
	}
	return first.End
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
