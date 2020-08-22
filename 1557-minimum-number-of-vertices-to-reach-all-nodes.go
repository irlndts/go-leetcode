package leetcode

type Item struct {
	ID      int
	To      []*Item
	Visited bool
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	m := make(map[int]*Item)

	for _, e := range edges {
		to, ok := m[e[1]]
		if !ok {
			to = &Item{ID: e[1]}
		}

		from, ok := m[e[0]]
		if !ok {
			from = &Item{ID: e[0]}
		}
		from.To = append(from.To, to)

		m[e[0]] = from
		m[e[1]] = to
	}

	for _, v := range m {
		for _, to := range v.To {
			path(to)
		}
	}

	result := []int{}
	for _, v := range m {
		if !v.Visited {
			result = append(result, v.ID)
		}
	}

	return result
}

func path(node *Item) {
	if node.Visited {
		return
	}
	node.Visited = true
	for _, to := range node.To {
		path(to)
	}
}
