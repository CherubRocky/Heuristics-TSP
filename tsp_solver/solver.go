package tsp_solver

import (
	"github.com/CherubRocky/Heuristics-TSP/simulated_annealing"
)

func (i *Instance) Solve(ids []int) simulated_annealing.Solution {
	initialSol := TravelSolution{Permutation: ids, instance: i}
	initSolCost := i.Cost(&initialSol)
	initialSol.SolutionCost = initSolCost
	ann := simulated_annealing.Annealer{
		Epsilon: 0.00005,
		ProbEps: 0.05,
		Temperature: 8.0,
		ColdingFactor: 0.9999,
		BatchSize: 2000,
		Best: &initialSol,
	}
	ann.InitialTemp(&initialSol, 0.87)
	return ann.ThresholdAcceptance(&initialSol)
}
