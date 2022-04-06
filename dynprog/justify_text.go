package dynprog

import (
	"fmt"
	"math"
	"strings"
)

type badnessAndSplit struct {
	badness int64
	split   int
}

type lineSplitter struct {
	words     []string
	cache     []badnessAndSplit
	lineWidth int
}

func (s *lineSplitter) nextSplit(i int) int64 {
	const debug = false
	badnessOf := func(i, j int) int64 {
		width := j - i - 1
		if width < 0 {
			width = 0
		}
		for _, w := range s.words[i:j] {
			width += len(w)
		}
		if width > s.lineWidth {
			return math.MaxInt32
		}
		t := int64(s.lineWidth - width)
		return t * t * t
	}

	if i == len(s.words) {
		return math.MaxInt32
	}

	best := len(s.words)
	bestBadness := badnessOf(i, len(s.words))
	for j := i + 1; j < len(s.words); j++ {
		badness := badnessOf(i, j)
		totalBadness := s.nextSplit(j) + badness
		if debug {
			fmt.Println("badness ", badness, " totalBadness ", totalBadness)
		}
		if totalBadness < bestBadness {
			best = j
			bestBadness = totalBadness

			if debug {
				fmt.Println("revise ", i, j)
			}
		}
	}
	if debug {
		fmt.Println(i, best)
	}
	s.cache[i] = badnessAndSplit{
		badness: bestBadness,
		split:   best,
	}
	return bestBadness
}

func SplitIntoLines(text string, lineWidth int) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return nil
	}
	splitter := lineSplitter{
		words:     words,
		lineWidth: lineWidth,
		cache:     make([]badnessAndSplit, len(words)),
	}
	splitter.nextSplit(0)
	var lines []string
	next := 0
	for next < len(words) {
		t := splitter.cache[next].split
		line := strings.Join(words[next:t], " ")
		lines = append(lines, line)
		next = t
	}
	return lines
}
