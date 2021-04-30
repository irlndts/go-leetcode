package problems

func findTheDifference(s string, t string) byte {
	m := map[rune]int{}
	for _, ss := range s {
		m[ss]++
	}

	for _, tt := range t {
		if v, ok := m[tt]; !ok || v == 0 {
			return byte(tt)
		}
		m[tt]--
	}

	var result byte
	return result
}
