package problems

// https://leetcode.com/problems/cheapest-flights-within-k-stops/

type Node struct {
	ID  int
	Dst []Prev

	Visited bool
	Price   int
}

type Prev struct {
	ID   int
	Cost int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	m := make(map[int]*Node)
	for i := 0; i < n; i++ {
		node := &Node{
			ID:  i,
			Dst: []Prev{},
		}
		m[i] = node
	}

	for _, flight := range flights {
		node := m[flight[1]]
		prev := m[flight[0]]
		price := flight[2]
		node.Dst = append(node.Dst, Prev{
			prev.ID,
			price,
		})
	}

	path(m, m[dst], src, 0, K, 0)

	if m[src].Price == 0 {
		return -1
	}

	return m[src].Price
}

func path(m map[int]*Node, current *Node, src, k, K, price int) {
	if current.ID == src {
		if !current.Visited || price < current.Price {
			current.Price = price
		}
		current.Visited = true
		return
	}

	if current.Visited && price >= current.Price {
		return
	}

	if k > K {
		return
	}

	current.Visited = true
	current.Price = price

	for _, next := range current.Dst {
		path(m, m[next.ID], src, k+1, K, next.Cost+price)
	}

}
