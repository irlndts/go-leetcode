package problems

import "strconv"

func permuteUnique(nums []int) [][]int {
	results = [][]int{}
	m = map[string]struct{}{}
	ar := make([]int, 0)
	helper(nums, ar)
	return results
}

var results [][]int
var m map[string]struct{}

func helper(nums, current []int) {
	if len(nums) == 0 {
		k := key(current)
		if _, ok := m[k]; ok {
			return
		}
		results = append(results, current)
		m[k] = struct{}{}
		return
	}

	for i, n := range nums {
		c := make([]int, len(nums))
		copy(c, nums)
		newCurrent := make([]int, len(current)+1)
		copy(newCurrent, current)
		newCurrent[len(newCurrent)-1] = n
		helper(append(c[:i], c[i+1:]...), newCurrent)
	}
}

func key(ar []int) string {
	result := ""
	for _, n := range ar {
		result += strconv.Itoa(n) + ":"
	}
	return result
}
