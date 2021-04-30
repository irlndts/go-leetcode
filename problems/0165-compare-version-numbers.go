package problems

import (
	"strconv"
	"strings"
)

func compareVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	vv1 := make([]int, 0, len(v1))
	for _, v := range v1 {
		n, _ := strconv.Atoi(v)
		vv1 = append(vv1, n)
	}

	vv2 := make([]int, 0, len(v2))
	for _, v := range v2 {
		n, _ := strconv.Atoi(v)
		vv2 = append(vv2, n)
	}

	for i := 0; i < max(len(vv1), len(vv2)); i++ {
		if i >= len(vv2) {
			return helper(vv1, i)
		}
		if i >= len(vv1) {
			return helper(vv2, i) * -1
		}

		if vv1[i] > vv2[i] {
			return 1
		}
		if vv2[i] > vv1[i] {
			return -1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(a []int, i int) int {
	for j := i; j < len(a); j++ {
		if a[j] != 0 {
			return 1
		}
	}
	return 0
}
