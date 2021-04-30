package problems

// https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
	a := 0
	for _, n := range nums {
		a ^= n
	}
	return a
}

func singleNumberHash(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := m[n]; ok {
			delete(m, n)
			continue
		}
		m[n] = struct{}{}
	}

	for k := range m {
		return k
	}
	return 0
}
