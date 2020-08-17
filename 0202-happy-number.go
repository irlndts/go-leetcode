package leetcode

func isHappy(N int) bool {
	found := make(map[int]struct{})
	for N != 0 {
		if _, ok := found[N]; ok {
			return false
		}
		found[N] = struct{}{}
		result := 0

		for N != 0 {
			tmp := N % 10
			result += tmp * tmp
			N = N / 10
		}
		N = result
		if N == 1 {
			return true
		}
	}

	return false
}
