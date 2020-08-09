package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	for _, n := range nums2 {
		if i == len(nums1) {
			break
		}
		for n > nums1[i] && i < m {
			i++
			if i == len(nums1) {
				return
			}
		}

		nums1 = append(nums1[:i+1], nums1[i:len(nums1)-1]...)
		nums1[i] = n
		i++
		m++
	}
}
