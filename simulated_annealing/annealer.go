package simulated_annealing

import "math"

type Annealer struct {
	Epsilon float64
	Temperature float64
	ColdingFactor float64
	BatchSize int
	Best Solution
}

func (a *Annealer) computeBatch(s Solution) (float64, Solution) {
	counter := 0
	var sum float64 = 0.0
	for counter < a.BatchSize {
		neigh := s.GetNeighbour()
		if neigh.Cost() <= s.Cost() + a.Temperature {
			if neigh.Cost() < a.Best.Cost() {
				a.Best = neigh
			}
			s = neigh
			counter++
			sum += s.Cost()
		}
	}
	return sum / float64(a.BatchSize), s
	
}

func (a *Annealer) ThresholdAcceptance(s Solution) Solution {
	p := 0.0
	for a.Temperature > a.Epsilon {
		q := math.MaxFloat64
		for p <= q {
			q = p
			p, s = a.computeBatch(s)
		}
		a.Temperature = a.ColdingFactor * a.Temperature
	}
	return a.Best
}
