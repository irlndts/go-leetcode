package problems

// https://leetcode.com/problems/count-and-say/

import "strconv"

func countAndSay(n int) string {
	result := "1"

	for i := 1; i < n; i++ {
		buf := ""
		var cur rune
		f := false
		m := make(map[rune]int)
		for _, r := range result {
			if f && cur != r {
				buf += strconv.Itoa(m[cur]) + string(cur)
				m[cur] = 0
			}
			f = true
			m[r]++
			cur = r
		}
		result = buf + strconv.Itoa(m[cur]) + string(cur)
	}

	return result
}
