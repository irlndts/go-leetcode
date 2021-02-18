package leetcode

import "sort"

// https://leetcode.com/problems/maximum-units-on-a-truck/

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	result := 0
	for _, box := range boxTypes {
		for i := 0; i < box[0]; i++ {
			if truckSize == 0 {
				return result
			}
			result += box[1]
			truckSize--
		}
	}
	return result
}
