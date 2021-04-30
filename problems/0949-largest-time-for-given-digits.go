package problems

import (
	"fmt"
	"strings"
)

func largestTimeFromDigits(A []int) string {
	//left := 23
	//right := 59

	var arr []string
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(A); k++ {
				if k == j || k == i {
					continue
				}
				for m := 0; m < len(A); m++ {
					if m == i || m == j || m == k {
						continue
					}
					arr = append(arr, fmt.Sprintf("%d%d:%d%d", A[i], A[j], A[k], A[m]))
				}
			}
		}
	}

	result := ""
	for _, a := range arr {
		parts := strings.Split(a, ":")
		if parts[0] > "23" {
			continue
		}
		if parts[1] > "59" {
			continue
		}
		if a > result {
			result = a
		}
	}

	return result
}
