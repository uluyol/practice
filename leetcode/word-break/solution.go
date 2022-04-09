package leetcode

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}

	suffixesValid := make([]bool, len(s))
Outer:
	for i := len(s) - 1; i >= 0; i-- {
		if words[s[i:]] {
			suffixesValid[i] = true
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if suffixesValid[j] {
				if words[s[i:j]] {
					suffixesValid[i] = true
					continue Outer
				}
			}
		}
	}

	return suffixesValid[0]
}
