package leetcode

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	if words[0] == "" {
		result := make([]int, 0, len(s))
		for i := 0; i <= len(s); i++ {
			result = append(result, i)
		}
		return result
	}

	total := len(words)

	m := mm(words)
	var results []int
	for i := 0; i <= len(s)-len(words[0]); i++ {
		word := s[i : i+len(words[0])]
		if _, ok := m[word]; ok {
			m[word]--
			total--
			if m[word] == 0 {
				delete(m, word)
			}
			if len(m) == 0 {
				results = append(results, i+len(word)-len(word)*len(words))
				m = mm(words)
				total = len(words)
				i = i - len(word)*(len(words)-1)
				continue
			}
			i = i + len(word) - 1
		} else if len(words) != total {
			i = i - len(word)*(len(words)-total)
			m = mm(words)
			total = len(words)
		}
	}

	return results
}

func mm(words []string) map[string]int {
	m := map[string]int{}
	for _, w := range words {
		m[w]++
	}
	return m
}
