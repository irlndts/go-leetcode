package leetcode

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/hamming-distance/

func hammingDistance(x int, y int) int {
	diff := x ^ y
	bits := 0
	for diff > 0 {
		if diff&1 != 0 {
			bits++
		}
		diff = diff >> 1
	}
	return bits
}

func hammingDistance1(x int, y int) int {
	a := fmt.Sprintf("%b", x)
	b := fmt.Sprintf("%b", y)

	a, b = foo(a, b)

	var result int
	for i := range a {
		if a[i] != b[i] {
			result++
		}
	}

	return result
}

func foo(a, b string) (string, string) {

	l := len(a)
	if len(b) > len(a) {
		l = len(b)
	}

	if len(a) < l {
		a = strings.Repeat("0", l-len(a)) + a
	}

	if len(b) < l {
		b = strings.Repeat("0", l-len(b)) + b
	}

	return a, b
}
