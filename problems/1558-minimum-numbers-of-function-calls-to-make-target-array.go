package problems

func minOperations(nums []int) int {
	result := 0
	found := false

	allNull := false
	for !allNull {
		allNull = true
		for i, n := range nums {
			if n%2 == 1 {
				nums[i]--
				result++
				found = true
			}
			if n != 0 {
				allNull = false
			}
		}
		if allNull {
			break
		}
		if !found {
			for i := range nums {
				nums[i] /= 2
			}
			result++
		}
		found = false
	}

	return result
}
