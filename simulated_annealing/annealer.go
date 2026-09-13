package simulated_annealing

import (
	"math"
	"fmt"
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
	for counter < a.BatchSize {
		neigh := s.GetNeighbour()
		if neigh.Cost() <= s.Cost() + a.Temperature {
			if neigh.Cost() < a.Best.Cost() {
				a.Best = neigh
			}
			fmt.Println("Accepted: ", neigh.Cost())
			s = neigh
			counter++
			sum += s.Cost()
		}
	}
	fmt.Println("Batch finished. Temperature: ", a.Temperature)
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
		neigh := s.GetNeighbour()
		if neigh.Cost() <= s.Cost() + t {
			if neigh.Cost() < a.Best.Cost() {
				a.Best = neigh
			}
			counter++
			s = neigh
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

/*func (a *Annealer) ThresholdAcceptance(s Solution) Solution {
	if a.Best == nil {
		a.Best = s
	}

	for a.Temperature > a.Epsilon {
		_, s = a.computeBatch(s)
		a.Temperature = a.ColdingFactor * a.Temperature
	}
	return a.Best
        }*/
