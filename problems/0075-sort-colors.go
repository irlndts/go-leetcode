package problems

func sortColors(nums []int) {
	m := map[int]int{
		0: 0,
		1: 0,
		2: 0,
	}

	for _, n := range nums {
		m[n]++
	}

	c := 0
	for i := range []int{0, 1, 2} {
		for j := 0; j < m[i]; j++ {
			nums[j+c] = i
		}
		c += m[i]
	}
}
