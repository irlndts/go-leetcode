package problems

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	var result int
	if len(prices) <= 1 {
		return result
	}

	l, r := prices[0], 0
	for _, p := range prices {
		r = p
		if r-l > result {
			result = r - l
		}

		if p < l {
			l = p
		}
	}
	return result
}
