package tsp

func Cost(permutation []int) float32 {
	var sum float32
	prev := 0
	for i, id := range permutation {
		if i > 0 {
			sum += GetDistance(prev, id)
		}
		prev = id
	}
	return sum
}


func CostNeighbour(permutation []int, swap1 int, swap2 int) float32 {
	
}

func GetDistance(int id1, int id2) float32 {

}