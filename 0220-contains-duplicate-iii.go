package leetcode

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k; j++ {
			if j >= len(nums) {
				break
			}
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
