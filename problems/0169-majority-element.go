package problems

// https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
