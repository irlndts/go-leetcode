package problems

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	arr := make([]string, 0, len(nums))
	for _, n := range nums {
		arr = append(arr, strconv.Itoa(n))
	}

	sort.Sort(ByLeftDigits(arr))
	result := ""
	for _, ar := range arr {
		result += ar
	}

	if result[0] == '0' {
		return "0"
	}

	return result
}

// ByLeftDigits sorts array by left digits
// i.e. 5 > 44 > 333
type ByLeftDigits []string

func (l ByLeftDigits) Len() int      { return len(l) }
func (l ByLeftDigits) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l ByLeftDigits) Less(i, j int) bool {
	return l[i]+l[j] > l[j]+l[i]
}
