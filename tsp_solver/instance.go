package tsp_solver
import (
	"math/rand/v2"
	
	"github.com/CherubRocky/Heuristics-TSP/simulated_annealing"
	"github.com/CherubRocky/Heuristics-TSP/tsp"
)


type Instance struct {
	Matrix *tsp.Cities
	Normalizer float64
	Random *rand.Rand
}


func (i *Instance) Cost(sol simulated_annealing.Solution) float64 {
	var sum float64
	tSol := sol.(*TravelSolution)
	prev := 0
	for index, id := range tSol.Permutation {
		if index > 0 {
			sum += i.Matrix.GetDistance(prev, id)
		}
		prev = id
	}
	return sum / i.Matrix.KEdgesSum// Modified this line
}

// Refactoring pending
func (i *Instance) Neighbour(sol simulated_annealing.Solution, swap1 int, swap2 int) simulated_annealing.Solution {
	tSol := sol.(*TravelSolution)
	max, min := getMaxAndMin(swap1, swap2)
	var maxEdges, minEdges, newEdges1, newEdges2 float64
	if max == min + 1 {
		if min > 0 {
			minEdges = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min - 1])
		}
		if max < len(tSol.Permutation) - 1 {
			maxEdges = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max + 1])
		}
		swap(max, min, tSol.Permutation)
		if min > 0 {
			newEdges1 = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min - 1])
		}
		if max < len(tSol.Permutation) - 1 {
			newEdges2 = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max + 1])
		}
		
	} else {
		if max != len(tSol.Permutation) - 1 && min != 0 {
			maxEdges = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max + 1]) + i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max - 1])
			minEdges = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min + 1]) + i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min - 1])
			swap(max, min, tSol.Permutation)
			newEdges1 = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max + 1]) + i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max - 1])
			newEdges2 = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min + 1]) + i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min - 1])
		} else if max == len(tSol.Permutation) - 1 && min != 0 {
			maxEdges = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max - 1])
			minEdges = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min + 1]) + i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min - 1])
			swap(max, min, tSol.Permutation)
			newEdges1 = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max - 1])
			newEdges2 = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min + 1]) + i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min - 1])
		} else if max != len(tSol.Permutation) - 1 && min == 0 {
			maxEdges = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max + 1]) + i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max - 1])
			minEdges = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min + 1])
			swap(max, min, tSol.Permutation)
			newEdges1 = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max + 1]) + i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max - 1])
			newEdges2 = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min + 1])
		} else if max == len(tSol.Permutation) - 1 && min == 0 {
			maxEdges = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max - 1])
			minEdges = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min + 1])
			swap(max, min, tSol.Permutation)
			newEdges1 = i.Matrix.GetDistance(tSol.Permutation[max], tSol.Permutation[max - 1])
			newEdges2 = i.Matrix.GetDistance(tSol.Permutation[min], tSol.Permutation[min + 1])
		}
	}
	balance := (-(maxEdges + minEdges) + (newEdges1 + newEdges2)) / i.Normalizer
	tSol.SolutionCost += balance
	tSol.lastDiff = balance
	tSol.swapInd1 = max
	tSol.swapInd2 = min
	return tSol
}

func (i *Instance) ComputeNormalizer() {}

func getMaxAndMin(a int, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

func swap(a int, b int, sol []int) {
	tmp := sol[a]
	sol[a] = sol[b]
	sol[b] = tmp
}
