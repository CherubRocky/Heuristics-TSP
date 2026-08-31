package simulated_annealing

type ProblemAdapter interface {
	Cost() float64
}
