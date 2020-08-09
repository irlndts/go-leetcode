package leetcode

import "strconv"

// https://leetcode.com/problems/fizz-buzz/

func fizzBuzz(n int) []string {
	results := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			results = append(results, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			results = append(results, "Fizz")
			continue
		}
		if i%5 == 0 {
			results = append(results, "Buzz")
			continue
		}
		results = append(results, strconv.Itoa(i))
	}
	return results
}
