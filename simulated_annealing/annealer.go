package simulated_annealing

import (
	"math"
	//	"fmt"
)

type Annealer struct {
	Epsilon float64
	ProbEps float64
	Temperature float64
	ColdingFactor float64
	BatchSize int
	Best Solution
}

func (a *Annealer) computeBatch(s Solution) (float64, Solution) {
	var sum float64 = 0.0
	counter := 0
	scape := 0
	for counter < a.BatchSize && scape < 25 * a.BatchSize {
		ogCost := s.Cost()
		neigh := s.GetNeighbour()
		nCost := neigh.Cost()
		if nCost <= ogCost + a.Temperature {
			if nCost < a.Best.Cost() {
				a.Best = neigh.Clone()
			}
			// fmt.Println("Accepted: ", neigh.Cost())
			counter++
			sum += s.Cost()
		} else {
			s.RollBack()
		}
		scape++
	}
	// fmt.Println("Batch finished. Temperature: ", a.Temperature)
	if counter == 0 {
		return s.Cost(), s
	}
	return sum / float64(counter), s
	
}

func (a *Annealer) ThresholdAcceptance(s Solution) Solution {
	p := 0.0
	for a.Temperature > a.Epsilon {
		q := math.MaxFloat64
		for {
			p, s = a.computeBatch(s)
			if q - p < a.Epsilon {
				break
			}
			q = p
		}
		a.Temperature = a.ColdingFactor * a.Temperature
	}
	return a.Best
}

func (a *Annealer) InitialTemp(s Solution, p float64) {
	var t1, t2 float64
	prob := a.AcceptedPercent(s, a.Temperature)
	if math.Abs(p - prob) <= a.ProbEps {
		return
	}
	if prob < p {
		for prob < p {
			a.Temperature *= 2
			prob = a.AcceptedPercent(s, a.Temperature)
		}
		t1 = a.Temperature / 2
		t2 = a.Temperature
	} else {
		for prob > p {
			a.Temperature = a.Temperature / 2
			prob = a.AcceptedPercent(s, a.Temperature)
		}
		t1 = a.Temperature
		t2 = a.Temperature * 2
	}
	a.Temperature = a.BinarySearch(s, t1, t2, p)
}

func (a *Annealer) AcceptedPercent(s Solution, t float64) float64 {
	counter := 0
	n := a.BatchSize
	for i := 0; i < n; i++ {
		ogCost := s.Cost()
		neigh := s.GetNeighbour()
		nCost := s.Cost()
		if nCost <= ogCost + t {
			if nCost < a.Best.Cost() {
				a.Best = neigh.Clone()
			}
			counter++
		} else {
			s.RollBack()
		}
	}
	return float64(counter) / float64(n)
}

func (a *Annealer) BinarySearch(s Solution, t1 float64, t2 float64, p float64) float64 {
	meanT := (t1 + t2) / 2
	if t2 - t1 < a.ProbEps {
		return meanT
	}
	prob := a.AcceptedPercent(s, meanT)
	if math.Abs(p - prob) < a.ProbEps {
		return meanT
	}
	if prob > p {
		return a.BinarySearch(s, t1, meanT, p)
	} else {
		return a.BinarySearch(s, meanT, t2, p)
	}
	
}
