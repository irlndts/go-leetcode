package leetcode

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}

	for i := 1; i <= len(s)/2; i++ {
		dif := len(s) / i
		if dif*i != len(s) {
			continue
		}
		h := hash(s[:i])
		for j := i; j < len(s); j += i {
			if h != hash(s[j:j+i]) {
				break
			}
			if j+i == len(s) {
				return true
			}
		}
	}

	return false
}

func hash(line string) int {
	a := 7
	m := 7919
	result := 0

	p := 1
	for i := len(line) - 1; i >= 0; i-- {
		result = (result + int(line[i])*p) % m
		p = (p * a) % m
	}

	return result
}
