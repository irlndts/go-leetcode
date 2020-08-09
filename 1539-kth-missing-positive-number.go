package leetcode

import "sort"

func findKthPositive(arr []int, k int) int {
	if len(arr) == 0 {
		return k
	}

	sort.Ints(arr)
	counter := 0
	number := 1
	for i := 0; i < len(arr); i++ {
		if arr[i] != number {
			counter++
			i--
		}

		if counter == k {
			return number
		}
		number++
	}

	return arr[len(arr)-1] + k - counter
}
