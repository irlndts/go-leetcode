package leetcode

import "sort"

func findRightInterval(intervals [][]int) []int {
	results := make([]int, 0, len(intervals))
	starts := map[int]int{}
	keys := []int{}
	for i, interval := range intervals {
		if _, ok := starts[interval[0]]; !ok {
			starts[interval[0]] = i
			keys = append(keys, interval[0])
		}
	}

	sort.Ints(keys)

	for _, interval := range intervals {
		if v, ok := starts[interval[1]]; ok {
			results = append(results, v)
			continue
		}

		found := false
		i := sort.Search(len(keys), func(i int) bool { return keys[i] >= interval[0] })
		for ; i < len(keys); i++ {
			if interval[1] <= keys[i] {
				results = append(results, starts[keys[i]])
				found = true
				break
			}
		}
		if found {
			continue
		}

		results = append(results, -1)
	}

	return results
}
