package leetcode

func coinChange(coins []int, amount int) int {
	cache := make(map[int]int)
	return coinChangeImpl(coins, amount, cache)
}

func coinChangeImpl(coins []int, amount int, cache map[int]int) int {
	// defer func() {
	// 	fmt.Printf("amount %d got %d\n", amount, cache[amount])
	// }()
	if n, ok := cache[amount]; ok {
		return n
	}
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	best := -1
	for _, c := range coins {
		n := coinChangeImpl(coins, amount-c, cache)
		if n < 0 {
			continue
		}
		n++ // add a coin for c
		if best < 0 || n < best {
			best = n
		}
	}

	cache[amount] = best
	return best
}
