package leetcode

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

var lettersByDigits = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var results []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	results = []string{}
	backtrack("", digits)
	return results
}

func backtrack(current, next string) {
	if len(next) == 0 {
		results = append(results, current)
		return
	}

	letters := lettersByDigits[rune(next[0])]
	for _, l := range letters {
		tmp := current + string(l)
		backtrack(tmp, next[1:])
	}
}
