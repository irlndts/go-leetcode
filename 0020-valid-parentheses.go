package leetcode

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	isRight := func(l string) bool {
		switch l {
		case ")", "}", "]":
			return true
		}
		return false
	}

	pop := func(arr []string) (string, []string) {
		if len(arr) == 0 {
			return "", nil
		}
		return arr[len(arr)-1], arr[:len(arr)-1]
	}

	isPair := func(left, right string) bool {
		pairs := map[string]string{
			"{": "}",
			"(": ")",
			"[": "]",
		}
		return pairs[left] == right
	}

	var lefts []string
	for _, right := range s {
		if isRight(string(right)) {
			var left string
			left, lefts = pop(lefts)
			if left == "" {
				return false
			}
			if !isPair(left, string(right)) {
				return false
			}
			continue
		}
		lefts = append(lefts, string(right))
	}
	return len(lefts) == 0
}
