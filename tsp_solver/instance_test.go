package tsp_solver
import (
	"github.com/CherubRocky/Heuristics-TSP/models"
	"github.com/CherubRocky/Heuristics-TSP/tsp"
	"math/rand/v2"
	
	"testing"
	"math"
)

var tsp150 = [150]int{1,2,3,4,5,6,7,8,9,11,12,14,16,17,19,20,22,23,25,26,27,74,75,77,163,164,165,166,167,168,169,171,172,173,174,176,179,181,182,183,184,185,186,187,244,297,326,327,328,329,330,331,332,333,334,336,339,340,343,344,345,346,347,349,350,351,352,353,444,483,489,490,491,492,493,494,495,496,499,500,501,502,504,505,507,508,509,510,511,512,520,652,653,654,655,656,657,658,660,661,662,663,665,666,667,668,670,671,673,674,675,676,678,815,816,817,818,819,820,821,822,823,825,826,828,829,832,837,839,840,978,979,980,981,982,984,985,986,988,990,991,995,999,1001,1003,1004,1037,1038,1073,1075}

var tsp40 = [40]int{1,2,3,4,5,6,7,75,163,164,165,168,172,244,327,329,331,332,333,489,490,491,492,493,496,652,653,654,656,657,815,816,817,820,978,979,980,981,982,984}

var best = [150]int{981,333,990,353,27,3,165,988,23,176,668,352,978,5,6,1004,351,22,676,665,344,490,653,507,2,656,667,184,832,661,663,657,829,1,9,986,508,168,505,19,496,839,815,173,673,182,172,163,329,509,493,979,837,995,331,662,8,999,674,504,511,327,670,350,840,336,980,297,244,339,502,1038,12,828,822,166,1037,494,167,495,328,326,169,1073,330,819,655,818,666,658,74,1001,340,675,186,520,16,671,179,512,821,1075,483,652,171,346,75,183,77,489,25,492,501,17,444,826,11,164,334,985,500,825,20,660,510,343,1003,349,984,491,499,347,817,4,174,991,185,26,820,654,7,823,678,816,187,982,14,181,345,332}

type testCase struct {
	description string
	ts *TravelSolution
	want float64
}

func TestCost(t *testing.T) {
	cases := []testCase {
		{"TSP 40", setupTestTS(tsp40[:], t), 4037072.076285812},
			{"TSP 150", setupTestTS(tsp150[:], t), 6092371.482090380},
			{"Best", setupTestTS(best[:], t), 0.162839},
	}
	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			costHelper(test, t)
			neighbourDeltaHelper(test, t)
		})
	}
}


func costHelper(tc testCase, t *testing.T) {
	t.Helper()
	ts := tc.ts
	instance := ts.instance

	got := instance.Cost(ts)
	want := tc.want
	margin := math.Abs(got - want)
	if margin > 0.000001 {
		t.Errorf("got %f, want %f", got, want)
	}
}

func neighbourDeltaHelper(tc testCase, t *testing.T) {
	t.Helper()
	ts := tc.ts
	instance := ts.instance
	
	ts.SolutionCost = instance.Cost(ts)

	for i := 0; i < 1000; i++ {
		u := instance.Random.IntN(len(ts.Permutation))
		v := instance.Random.IntN(len(ts.Permutation))
		for v == u {
			v = instance.Random.IntN(len(ts.Permutation))
		}

		neigh := instance.Neighbour(ts, u, v)

		realCost := instance.Cost(neigh)

		margin := math.Abs(neigh.Cost() - realCost)
		if margin > 0.000001 {
			t.Fatalf("swap(%d, %d) failed: Delta gave %f, but Cost gave %f", u, v, neigh.Cost(), realCost)
		}
	}
}

func setupTestTS(arr []int, t *testing.T) *TravelSolution {
	t.Helper()

	datBase, err := models.NewDB()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	citsSlice, err := datBase.QueryCities(arr[:])
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	conns, err := datBase.QueryConnections(arr[:])
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	err = datBase.Close()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	cities := tsp.NewCities(conns, citsSlice)

	instance := Instance{Matrix: cities, Normalizer: cities.KEdgesSum, Random: rand.New(rand.NewPCG(1, 2)),}

	ts := TravelSolution{
		Permutation:  arr[:],
		SolutionCost: 0.0,
		instance:     &instance,
	}
	return &ts
}
