package leetcode

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	var result [][]string

	var i int
	m := make(map[string]int)
	for _, str := range strs {
		ls := strings.Split(str, "")
		sort.Strings(ls)
		key := strings.Join(ls, "")
		v, ok := m[key]
		if !ok {
			m[key] = i
			result = append(result, []string{str})
			i++
			continue
		}

		result[v] = append(result[v], str)
	}

	return result
}
