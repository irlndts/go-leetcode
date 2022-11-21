package problems

// https://leetcode.com/problems/intersection-of-two-arrays/

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, num := range nums1 {
		m[num] = struct{}{}
	}

	m2 := make(map[int]struct{})
	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			m2[num] = struct{}{}
		}
	}

	result := []int{}
	for k, _ := range m2 {
		result = append(result, k)
	}

	return result
}
