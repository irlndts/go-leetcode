package problems

// https://leetcode.com/problems/separate-the-digits-in-an-array/t st

func separateDigits(nums []int) []int {
	var result []int
	for i := len(nums) - 1; i >= 0; i-- {
		cur := nums[i]

		for cur != 0 {
			a := cur % 10
			result = append(result, a)
			cur = (cur - a) / 10
		}
	}

	res := make([]int, 0, len(result))
	for i := len(result) - 1; i >= 0; i-- {
		res = append(res, result[i])
	}

	return res
}
