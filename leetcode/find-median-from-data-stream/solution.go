package leetcode

type MedianFinder struct {
	t   RadixTree
	num int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.t.Add(kTopPos, num+kOffset)
	this.num++
}

func (this *MedianFinder) FindMedian() float64 {
	num := this.num
	if (num % 2) == 0 {
		a := this.t.FindIndex(kTopPos, num/2-1, 0, 0) - kOffset
		b := this.t.FindIndex(kTopPos, num/2, 0, 0) - kOffset
		// fmt.Println(a)
		// fmt.Println(b)
		return float64(b-a)/2 + float64(a)
	}
	return float64(this.t.FindIndex(kTopPos, num/2, 0, 0) - kOffset)
}

type RadixTree struct {
	buckets      [10]*RadixTree
	bucketCounts [10]int
}

func (t *RadixTree) Add(pos, n int) {
	for pos >= 1 {
		bkt := n / pos
		pos, n = pos/10, n%pos
		if pos > 0 && t.buckets[bkt] == nil {
			t.buckets[bkt] = new(RadixTree)
		}
		t.bucketCounts[bkt]++
		t = t.buckets[bkt]
	}
}

func (t *RadixTree) FindIndex(pos, index, cumIndex, partialSum int) int {
	for pos > 0 {
		for i := 0; i < 10; i++ {
			bktSize := t.bucketCounts[i]
			if bktSize > 0 && cumIndex+bktSize > index {
				// It is in buckets[i]
				partialSum += i * pos
				pos /= 10
				t = t.buckets[i]
				break
			}
			cumIndex += bktSize
		}
	}
	return partialSum
}

const kOffset = 100_000
const kTopPos = 1_000_000
