package leetcode

import (
	"fmt"
	"sort"
)

const debug = false

func eraseOverlapIntervals(intervals [][]int) int {
	finder := finder{
		_I: intervals,
	}
	finder.init()
	return len(intervals) - finder.maxConflictFree()
}

// finder find a maximum conflict-free subset (MCFS) in intervals (I)
//
// # Approach
//
// First, sort I by end times (then start times)
//
// {I[0]} is trivially an MCFS for I[:1]
// If I[1] is conflict-free with I[0], then an MCFS for I[:2] is {I[0], I[1]}
// If they do conflict, then both {I[0]} and {I[1]} are MCFS for I[:2].
//
// Define a smallest-end MCFS (SE-MCFS) as an MCFS with the earliest possible
// end time.
// In I[:2] above, {I[0]} would be the SE-MCFS for I[:2].
//
// In general, construct a SE-MCFS for I[:i] using a SE-MCFS for I[:i-1].
// If the SE-MCFS we have for I[:i-1] conflicts with I[i], then no MCFS can
// be larger than the SE-MCFS for I[:i-1] so use that.
// Otherwise, add I[i] to the SE-MCFS of I[:i-1].
type finder struct {
	_I [][]int // sorted
	// _I[pred[i]] is the last element in the SE-MCFS we've constructed
	pred []int
}

func (f *finder) maxConflictFree() int {
	if len(f._I) == 0 {
		return 0
	}
	f.pred[0] = 0 // _I[0] is always an SE-MCFS of _I[:1]
	size := 1
	for i := 1; i < len(f._I); i++ {
		prevPred := f.pred[i-1]
		last := f._I[prevPred]
		cur := f._I[i]
		if last[1] <= cur[0] {
			// No conflict
			f.pred[i] = i
			size++
		} else {
			f.pred[i] = prevPred
		}
	}
	return size
}

func (f *finder) init() {
	intervals := f._I
	// Sort in increasing order or start, then end time
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	f.pred = make([]int, len(intervals))
	if debug {
		fmt.Printf("(sorted) intervals = %v\n", f._I)
	}
}
