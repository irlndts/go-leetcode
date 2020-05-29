package leetcode

// https://leetcode.com/problems/combination-sum/

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combination(candidates, target)
}

func combination(candidates []int, target int) [][]int {
	results := [][]int{}

	for i, c := range candidates {
		if c == target {
			results = append(results, []int{c})
			continue
		}

		if c > target {
			break
		}

		tmps := combination(candidates[i:], target-c)
		for _, tmp := range tmps {
			tmp = append(tmp, c)
			results = append(results, tmp)
		}
	}

	return results
}
