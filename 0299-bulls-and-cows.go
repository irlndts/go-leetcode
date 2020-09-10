package leetcode

import "fmt"

func getHint(secret, guess string) string {
	m := make(map[rune]int)
	for _, s := range secret {
		m[s]++
	}

	a := 0
	b := 0
	for i := range secret {
		k := rune(guess[i])
		if secret[i] == guess[i] {
			a++
			m[k]--
		}
	}

	for i := range secret {
		k := rune(guess[i])
		if secret[i] != guess[i] {
			if _, ok := m[k]; ok {
				m[k]--
				if m[k] >= 0 {
					b++
				}
			}
		}
	}

	return fmt.Sprintf("%dA%dB", a, b)
}
