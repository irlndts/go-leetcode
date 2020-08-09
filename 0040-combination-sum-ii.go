package leetcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	sort.Ints(candidates)

	results = map[string]struct{}{}
	for i := 0; i < len(candidates); i++ {
		Solution(candidates, target, 0, i, "")
	}

	output := make([][]int, 0, len(results))
	for r := range results {
		numbers := strings.Split(r, ":")
		arr := make([]int, 0, len(numbers))
		for _, n := range numbers {
			number, _ := strconv.Atoi(n)
			arr = append(arr, number)
		}
		output = append(output, arr)
	}

	return output
}

var results map[string]struct{}

func Solution(candidates []int, target, sum, index int, key string) {
	sum += candidates[index]
	if sum > target {
		delete(results, key)
		return
	}

	if len(key) == 0 {
		key = strconv.Itoa(candidates[index])
	} else {
		key = fmt.Sprintf("%s:%d", key, candidates[index])
	}
	if sum == target {
		results[key] = struct{}{}
		return
	}

	for i := index + 1; i < len(candidates); i++ {
		Solution(candidates, target, sum, i, key)
	}
}
