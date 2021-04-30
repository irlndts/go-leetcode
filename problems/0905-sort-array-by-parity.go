package problems

func sortArrayByParity(A []int) []int {
	l := 0
	for i, num := range A {
		if num%2 == 0 {
			A[i], A[l] = A[l], A[i]
			l++
		}
	}
	return A
}
