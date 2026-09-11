package tsp_solver

import (
	"github.com/CherubRocky/Heuristics-TSP/simulated_annealing"

	"math/rand/v2"
)

type TravelSolution struct {
	Permutation []int
	SolutionCost float64
	instance *Instance
}

func (ts *TravelSolution) Cost() float64 {
	if ts.SolutionCost == 0.0 {
		ts.SolutionCost = ts.instance.Cost(ts)
	}
	return ts.SolutionCost
}

func (ts *TravelSolution) GetNeighbour() simulated_annealing.Solution {
	i, j := getIndexes(len(ts.Permutation), ts.instance.Random)
	return ts.instance.Neighbour(ts, i, j)
}

func getIndexes(l int, random *rand.Rand) (int, int) {
	var index1, index2 int
	index1 = random.IntN(l)
	index2 = random.IntN(l)
	for index1 == index2 {
		index2 = random.IntN(l)
	}
	return index1, index2
}
