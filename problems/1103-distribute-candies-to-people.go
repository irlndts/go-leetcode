package problems

func distributeCandies(candies int, numPeople int) []int {
	result := make([]int, numPeople)
	index := 0
	current := 1
	for candies-current >= 0 {
		result[index] += current
		candies -= current
		current++

		index++
		if index == numPeople {
			index = 0
		}
	}
	result[index] += candies
	return result
}
