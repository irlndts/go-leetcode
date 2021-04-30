package problems

var result [][]int

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	result = [][]int{}

	foo(nums, []int{})
	return result
}

func foo(nums, row []int) {
	if len(nums) == 0 {
		result = append(result, row)
		return
	}

	for i, n := range nums {
		ar := cpy(nums)
		ar[0], ar[i] = ar[i], ar[0]
		next := append(row, n)
		foo(ar[1:], next)
	}
}

func cpy(nums []int) []int {
	ar := make([]int, 0, len(nums))
	for _, n := range nums {
		ar = append(ar, n)
	}
	return ar
}
