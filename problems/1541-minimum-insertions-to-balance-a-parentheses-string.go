package problems

import "errors"

func minInsertions(s string) int {
	if s == "" {
		return 0
	}

	result := 0

	input := s

	queue := NewQueue()
	for _, r := range input {
		s := string(r)
		queue.Put(s)
	}

	counter := 0

	for {
		val, err := queue.Pop()
		if err != nil {
			break
		}
		switch val {
		case ")":
			next, err := queue.Pop()
			if err != nil {
				result++
				counter++
				break
			}
			if next != ")" {
				result++
				queue.Put(next)
			}
			counter++
		case "(":
			if counter == 0 {
				result += 2
				continue
			}
			counter--
		}
	}

	result += counter

	return result
}

// Node ...
type Node struct {
	Val   string
	Right *Node
	Left  *Node
}

// Queue ...
type Queue struct {
	Head *Node
	Tail *Node
}

// NewQueue ...
func NewQueue() *Queue {
	return &Queue{}
}

// Put ...
func (q *Queue) Put(val string) {
	if q.Head == nil {
		q.Head = &Node{
			Val: val,
		}
		q.Tail = q.Head
		return
	}

	q.Tail.Right = &Node{
		Val:  val,
		Left: q.Tail,
	}

	q.Tail = q.Tail.Right
}

// Get ...
func (q *Queue) Get() (string, error) {
	if q.Head == nil {
		return "", errors.New("error")
	}

	val := q.Head.Val
	if q.Head.Right == nil {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.Right
		q.Head.Left = nil
	}

	return val, nil
}

// Pop ...
func (q *Queue) Pop() (string, error) {
	if q.Tail == nil {
		return "", errors.New("error")
	}

	val := q.Tail.Val
	if q.Tail.Left == nil {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Tail = q.Tail.Left
		q.Tail.Right = nil
	}

	return val, nil

}

var pairs = map[string]string{
	"(": ")",
}
