package leetcode

func findAnagrams(s string, p string) []int {
	input, template := s, p
	if len(input) < len(template) {
		return []int{}
	}

	mTemplate := map[rune]int{}
	for _, t := range template {
		mTemplate[t]++
	}

	result := []int{}
	anagram := false
	for i := 0; i <= len(input)-len(template); i++ {
		if _, ok := mTemplate[rune(input[i])]; !ok {
			continue
		}

		if anagram && input[i-1] == input[i+len(template)-1] {
			result = append(result, i)
			continue
		}

		if _, ok := mTemplate[rune(input[i+len(template)-1])]; !ok {
			i += len(template) - 1
			continue
		}

		anagram = IsAnagram(input[i:i+len(template)], mTemplate)
		if anagram {
			result = append(result, i)
		}
	}
	return result
}

// IsAnagram ...
func IsAnagram(input string, template map[rune]int) bool {
	m := map[rune]int{}
	for _, r := range input {
		m[r]++
	}

	if len(m) != len(template) {
		return false
	}

	for k := range m {
		if m[k] != template[k] {
			return false
		}
	}

	return true
}
