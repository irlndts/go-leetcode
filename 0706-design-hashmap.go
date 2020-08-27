package leetcode

type MyHashMap struct {
	size  int
	table []List
}

func (m *MyHashMap) key(num int) int {
	return num % m.size
}

type Item struct {
	key   int
	val   int
	right *Item
}

type List struct {
	head *Item
}

func (l *List) put(key, val int) {
	if l.head == nil {
		l.head = &Item{key: key, val: val}
		return
	}

	// rewrite value if exists
	start := l.head
	for start != nil {
		if start.key == key {
			start.val = val
			return
		}
		start = start.right
	}

	item := &Item{key: key, val: val}
	item.right = l.head
	l.head = item
}

func (l *List) delete(key int) int {
	start := l.head
	var prev *Item
	for start != nil {
		if start.key == key {
			// remove element from list
			if prev != nil {
				prev.right = start.right
			} else {
				l.head = start.right
			}
			return start.val
		}

		prev = start
		start = start.right
	}

	// nothing was found, according to task return -1
	return -1
}

func (l *List) get(key int) int {
	start := l.head
	for start != nil {
		if start.key == key {
			return start.val
		}
		start = start.right
	}

	// nothing was found, according to task return -1
	return -1
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		size:  1009,
		table: make([]List, 1009),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.table[this.key(key)].put(key, value)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.table[this.key(key)].get(key)
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.table[this.key(key)].delete(key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
