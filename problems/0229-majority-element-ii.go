package problems

func majorityElement(nums []int) []int {
	m := map[int]int{}
	result := []int{}
	for _, num := range nums {
		m[num]++
		if m[num] == len(nums)/3+1 {
			result = append(result, num)
		}
	}
	return result
}
