package leetcode

// https://leetcode.com/problems/single-number-ii/

func singleNumber(nums []int) int {
	m := map[int]int{}

	for _, n := range nums {
		m[n]++
		if m[n] == 3 {
			delete(m, n)
		}
	}

	for k := range m {
		return k
	}

	return 0
}
