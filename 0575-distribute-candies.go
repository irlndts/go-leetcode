package leetcode

// https://leetcode.com/problems/distribute-candies/

func distributeCandies(candyType []int) int {
	l := map[int]struct{}{}
	h := len(candyType) / 2
	for _, candy := range candyType {
		l[candy] = struct{}{}
		if h < len(l) {
			return len(candyType) / 2
		}
	}

	return len(l)
}
