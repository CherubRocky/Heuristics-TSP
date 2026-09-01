package tsp_solver
import (
	"github.com/CherubRocky/Heuristics-TSP/simulated_annealing"
	"github.com/CherubRocky/Heuristics-TSP/tsp"
)


type Instance struct {
	Matrix *tsp.Cities
	Normalizer float64
}


func (i *Instance) Cost(sol simulated_annealing.Solution) float64 {
	return sol.Cost() // Implementation pending
}

func (i *Instance) CostNeighbour(sol simulated_annealing.Solution, swap1 int, swap2 int) float64 {
	return sol.Cost() // Implementation pending
}

func (i *Instance) ComputeNormalizer() {}
