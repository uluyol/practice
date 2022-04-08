package leetcode

import "sort"

func search(nums []int, target int) int {
	offset := findMinIndex(nums)
	id := sort.Search(len(nums), func(i int) bool {
		return target <= nums[(offset+i)%len(nums)]
	})
	realID := (offset + id) % len(nums)
	if 0 <= id && id < len(nums) && nums[realID] == target {
		return realID
	}
	return -1
}

func findMinIndex(nums []int) int {
	// Starting Pos: BBBBLLLLLLLL
	//
	// Step 1: Find a small window that also looks like this.
	// - Pick 0, mid, len-1
	// - Look for same pattern [0, mid) and [mid, len-1)
	// - Recurse into the one that has it
	// - If we can't find one, then the subslices are not rotated. min is at the start of one
	//
	// Step 2: Once we have a small enough slice, just do a linear search.
	// - A linear search is O(N), but N is small

	indexOff := 0
	const linearProbSize = 8
	for {
		if len(nums) < 8 {
			return indexOff + findMinIndexLinear(nums)
		}

		mid := len(nums) / 2
		numsA := nums[:mid]
		numsB := nums[mid:]

		if looksRotated(numsA) {
			// fmt.Printf("recurse A: beg: %d end: %d\n", numsA[0], numsA[len(numsA)-1])
			nums = numsA
			continue
		}
		if looksRotated(numsB) {
			// fmt.Printf("recurse B: beg: %d end: %d\n", numsB[0], numsB[len(numsB)-1])
			nums = numsB
			indexOff += mid
			continue
		}
		// Neither is rotated. So min is first entry in one of them
		// fmt.Println("not rotated")
		a := numsA[0]
		b := numsB[0]
		if a < b {
			return indexOff
		}
		return indexOff + mid
	}
}

func findMinIndexLinear(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	min := nums[0]
	minIndex := 0
	for i, v := range nums {
		if v < min {
			min = v
			minIndex = i
		}
	}
	return minIndex
}

func looksRotated(nums []int) bool {
	return nums[0] >= nums[len(nums)-1]
}
