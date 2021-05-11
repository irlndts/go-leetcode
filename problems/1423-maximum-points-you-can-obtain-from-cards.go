package problems

// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards

func maxScore(arr []int, k int) int {
	total := 0
	for _, v := range arr{
		total += v
	}

	current := 0
	for i := 0; i < k; i++{
		current+= arr[i]
	}

	max := current
	total -= current

	for i := 0; i < k; i++ {
		total += arr[k-i-1]
		total -= arr[len(arr)-i-1]

		current -= arr[k-i-1]
		current += arr[len(arr)-i-1]
		if current > max {
			max = current
		}
	}

	return max
}