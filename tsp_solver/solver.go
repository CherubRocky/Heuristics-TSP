package tsp_solver

import (
	"github.com/CherubRocky/Heuristics-TSP/simulated_annealing"
)

func (i *Instance) Solve(ids []int) simulated_annealing.Solution {
	initialSol := TravelSolution{Permutation: ids, instance: i}
	initSolCost := i.Cost(&initialSol)
	initialSol.SolutionCost = initSolCost
	ann := simulated_annealing.Annealer{
		Epsilon: 0.00001,
		Temperature: 0.5,
		ColdingFactor: 0.9,
		BatchSize: 2000,
		Best: &initialSol,
	}
	return ann.ThresholdAcceptance(&initialSol)
}
