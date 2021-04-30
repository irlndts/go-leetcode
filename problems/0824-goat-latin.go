package problems

import "strings"

var vovels = map[byte]struct{}{
	'a': struct{}{},
	'e': struct{}{},
	'i': struct{}{},
	'o': struct{}{},
	'u': struct{}{},
	'A': struct{}{},
	'E': struct{}{},
	'I': struct{}{},
	'O': struct{}{},
	'U': struct{}{},
}

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	for i, word := range words {
		if _, ok := vovels[word[0]]; ok {
			words[i] = word + "ma" + strings.Repeat("a", i+1)
		} else {
			words[i] = word[1:] + word[0:1] + "ma" + strings.Repeat("a", i+1)
		}
	}
	return strings.Join(words, " ")
}
