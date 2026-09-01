package simulated_annealing

type Solution interface {
	Cost() float64
	GetNeighbour() Solution
}
