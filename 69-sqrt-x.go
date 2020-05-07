package leetcode

// https://leetcode.com/problems/sqrtx/

func mySqrt(x int) int {
	l, r := 0, x

	for l <= r {
		m := l + (r-l)/2
		mm := m * m
		if mm == x {
			return m
		}
		if mm > x {
			r = m - 1
			continue
		}
		l = m + 1
	}

	return r
}
