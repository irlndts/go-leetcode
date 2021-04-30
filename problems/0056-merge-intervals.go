package problems

import "sort"

type Interval struct {
	Start int
	End   int
}

type ByEnd []Interval

func (b ByEnd) Len() int      { return len(b) }
func (b ByEnd) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByEnd) Less(i, j int) bool {
	if b[i].Start == b[j].Start {
		return b[i].End < b[j].End
	}
	return b[i].Start < b[j].Start
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	ii := make([]Interval, 0, len(intervals))
	for _, i := range intervals {
		ii = append(ii, Interval{i[0], i[1]})
	}

	sort.Sort(ByEnd(ii))

	var result [][]int
	cur := ii[0]
	for _, interval := range ii[1:] {
		if interval.Start <= cur.End {
			if cur.End < interval.End {
				cur.End = interval.End
			}
			continue
		}

		result = append(result, []int{cur.Start, cur.End})
		cur = interval
	}
	result = append(result, []int{cur.Start, cur.End})

	return result
}
