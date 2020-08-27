package leetcode

import "sort"

func twoCitySchedCost(costs [][]int) int {
	items := make([]Item, 0, len(costs))
	for _, cost := range costs {
		item := Item{cost[0], cost[1]}
		items = append(items, item)
	}

	sort.Sort(ByK(items))
	result := 0
	for i := 0; i < len(items)/2; i++ {
		result += items[i].l
	}
	for i := len(items) / 2; i < len(items); i++ {
		result += items[i].r
	}
	return result
}

type Item struct {
	l, r int
}

type ByK []Item

func (k ByK) Len() int      { return len(k) }
func (k ByK) Swap(i, j int) { k[i], k[j] = k[j], k[i] }
func (k ByK) Less(i, j int) bool {
	return (k[i].l - k[i].r) < (k[j].l - k[j].r)
}
