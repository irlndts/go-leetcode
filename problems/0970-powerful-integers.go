package problems

// https://leetcode.com/problems/powerful-integers/

func powerfulIntegers(x int, y int, bound int) []int {
	m := map[int]struct{}{}

	yHash := map[int]int{}
	xHash := collectPows(x, bound)
	if x == y {
		yHash = xHash
	} else {
		yHash = collectPows(y, bound)
	}

	for _, xx := range xHash {
		for _, yy := range yHash{
			if xx+yy > bound {
				continue
			}
			m[xx+yy] = struct{}{}
		}
	}

	result := make([]int, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	// sort.Ints(result)
	return result
}


func collectPows(a, bound int) map[int]int {
	if bound == 0 {
		return nil
	}
	hash := map[int]int{
		0: 1,
	}
	if a == 1 {
		return hash
	}

	prev := a
	for i := 0; i <= bound; i++{
		cur, ok := hash[i]
		if ok {
			prev = cur
			continue
		}
		cur = prev * a
		if cur > bound {
			return hash
		}
		hash[i] = cur
		prev = cur
	}
	return hash
}
