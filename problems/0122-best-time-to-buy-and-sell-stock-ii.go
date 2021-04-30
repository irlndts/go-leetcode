package problems

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	left, prev := prices[0], prices[0]
	result := 0
	for i := 1; i < len(prices); i++ {
		number := prices[i]
		if number < prev {
			result += prev - left
			left = number
		}
		prev = number
	}
	if prev-left > 0 {
		result += prev - left
	}
	return result
}
