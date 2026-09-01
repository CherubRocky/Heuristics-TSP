package tsp_solver

import (
	"github.com/CherubRocky/Heuristics-TSP/simulated_annealing"
)

type TravelSolution struct {
	Permutation []int
	SolutionCost float64
}

func (ts *TravelSolution) Cost() float64 {
	return 0
}

func (ts *TravelSolution) GetNeighbour() simulated_annealing.Solution {
	return &TravelSolution{}
}
