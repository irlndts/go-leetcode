package leetcode

func mincostTickets(days, costs []int) int {
	memo := make([]int, 366)

	set := make(map[int]struct{})
	for _, day := range days {
		set[day] = struct{}{}
	}

	return dp(memo, costs, set, 1)
}

func dp(memo, costs []int, set map[int]struct{}, day int) int {
	if day > 365 {
		return 0
	}
	if memo[day] != 0 {
		return memo[day]
	}

	ans := 0
	if _, ok := set[day]; ok {
		ans = min(dp(memo, costs, set, day+1)+costs[0], dp(memo, costs, set, day+7)+costs[1])
		ans = min(ans, dp(memo, costs, set, day+30)+costs[2])
	} else {
		ans = dp(memo, costs, set, day+1)
	}

	memo[day] = ans
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
