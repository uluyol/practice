package dynprog

import (
	"math"
)

type OpData[T any] interface {
	OpCost(o T) float64
	IfApplyOp(o T) T
}

func BestParenthesizeCost[S []T, T OpData[T]](ops S) float64 {
	opt := &costOptimizer[S, T]{
		ops,
		make([]float64, len(ops)*len(ops)+1),
		make([]T, len(ops)*len(ops)+1),
	}
	for i := range opt.bestCost {
		opt.bestCost[i] = -1
	}
	return opt.optimizeCost(0, len(ops))
}

type costOptimizer[S []T, T OpData[T]] struct {
	ops      S
	bestCost []float64
	bestData []T
}

func (o *costOptimizer[S, T]) bestCostSlot(i, j int) *float64 {
	return &o.bestCost[i*len(o.ops)+j]
}

func (o *costOptimizer[S, T]) setBestData(i, j int, val T) {
	o.bestData[i*len(o.ops)+j] = val
}

func (o *costOptimizer[S, T]) getBestData(i, j int) T {
	return o.bestData[i*len(o.ops)+j]
}

func (o *costOptimizer[S, T]) optimizeCost(i, j int) float64 {
	costSlot := o.bestCostSlot(i, j)
	if *costSlot >= 0 {
		return *costSlot
	}
	ops := o.ops[i:j]
	if len(ops) == 0 {
		panic("zero ops")
	} else if len(ops) == 1 {
		*costSlot = 0
		o.setBestData(i, j, ops[0])
		return *costSlot
	}
	var (
		lowestCost     = math.MaxFloat64
		lowestCostData T
	)
	for breakPoint := i + 1; breakPoint < j; breakPoint++ {
		cost := o.optimizeCost(i, breakPoint)
		cost += o.optimizeCost(breakPoint, j)
		d1 := o.getBestData(i, breakPoint)
		d2 := o.getBestData(breakPoint, j)
		cost += d1.OpCost(d2)
		// fmt.Printf("cost %g, a = %v b = %v\n", cost, d1, d2)
		if cost < lowestCost {
			lowestCost = cost
			lowestCostData = d1.IfApplyOp(d2)
		}
	}
	*costSlot = lowestCost
	o.setBestData(i, j, lowestCostData)
	return *costSlot
}
