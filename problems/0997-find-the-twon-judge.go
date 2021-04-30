package problems

func findJudge(N int, trust [][]int) int {
	if N == 1 && len(trust) == 0 {
		return N
	}

	trusted := make(map[int]int)
	trusts := make(map[int]bool)

	for _, t := range trust {
		trusted[t[1]]++
		trusts[t[0]] = true
	}

	for k, v := range trusted {
		if v == N-1 {
			if _, ok := trusts[k]; !ok {
				return k
			}
		}
	}

	return -1
}
