package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxEndingWith := make([]int, len(nums))
	maxEndingWith[0] = nums[0]
	for i := 1; i < len(maxEndingWith); i++ {
		v := nums[i]
		if maxEndingWith[i-1] > 0 {
			v += maxEndingWith[i-1]
		}
		maxEndingWith[i] = v
	}
	max := maxEndingWith[0]
	for i := 1; i < len(maxEndingWith); i++ {
		v := maxEndingWith[i]
		if v > max {
			max = v
		}
	}
	return max
}
