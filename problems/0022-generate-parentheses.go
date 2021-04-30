package problems

// https://leetcode.com/problems/generate-parentheses/

var results []string

func generateParenthesis(n int) []string {
	results = []string{}
	generate("", 0, 0, n)
	return results
}

func generate(current string, opened, closed, n int) {
	if len(current) == n*2 {
		results = append(results, current)
		return
	}

	if opened > closed {
		generate(current+")", opened, closed+1, n)
	}

	if opened != n {
		generate(current+"(", opened+1, closed, n)
	}
}
