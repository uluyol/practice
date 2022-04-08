package leetcode

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxEndingWith := make([]int, len(nums))
	maxNegEndingWith := make([]int, len(nums))

	maxEndingWith[0] = nums[0]
	maxNegEndingWith[0] = nums[0]

	for i := 1; i < len(maxNegEndingWith); i++ {
		v := nums[i]
		max := v
		maxNeg := v
		{
			v2 := maxEndingWith[i-1] * v
			if v2 > max {
				max = v2
			}
			if v2 < maxNeg {
				maxNeg = v2
			}
		}
		if maxNegEndingWith[i-1] < 0 {
			v2 := maxNegEndingWith[i-1] * v
			if v2 > max {
				max = v2
			}
			if v2 < maxNeg {
				maxNeg = v2
			}
		}

		maxEndingWith[i] = max
		maxNegEndingWith[i] = maxNeg
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
