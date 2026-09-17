package tsp_solver
import (
	"github.com/CherubRocky/Heuristics-TSP/models"
	"github.com/CherubRocky/Heuristics-TSP/tsp"
	"math/rand/v2"
	
	"testing"
	"math"
)

func TestCost(t *testing.T) {
	var tsp150 = [150]int{1,2,3,4,5,6,7,8,9,11,12,14,16,17,19,20,22,23,25,26,27,74,75,77,163,164,165,166,167,168,169,171,172,173,174,176,179,181,182,183,184,185,186,187,244,297,326,327,328,329,330,331,332,333,334,336,339,340,343,344,345,346,347,349,350,351,352,353,444,483,489,490,491,492,493,494,495,496,499,500,501,502,504,505,507,508,509,510,511,512,520,652,653,654,655,656,657,658,660,661,662,663,665,666,667,668,670,671,673,674,675,676,678,815,816,817,818,819,820,821,822,823,825,826,828,829,832,837,839,840,978,979,980,981,982,984,985,986,988,990,991,995,999,1001,1003,1004,1037,1038,1073,1075}
	datBase, err := models.NewDB()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	citsSlice, err := datBase.QueryCities(tsp150[:])
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	conns, err := datBase.QueryConnections(tsp150[:])
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	err = datBase.Close()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	cities := tsp.NewCities(conns, citsSlice)

	instance := Instance{Matrix: cities, Normalizer: cities.KEdgesSum,}
	
	ts := TravelSolution{Permutation: tsp150[:],}

	got := instance.Cost(&ts)
	want := 6092371.482090380
	margin := math.Abs(got - want)
	if margin > 0.000001 {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestNeighbourDelta(t *testing.T) {
		var tsp150 = [150]int{1,2,3,4,5,6,7,8,9,11,12,14,16,17,19,20,22,23,25,26,27,74,75,77,163,164,165,166,167,168,169,171,172,173,174,176,179,181,182,183,184,185,186,187,244,297,326,327,328,329,330,331,332,333,334,336,339,340,343,344,345,346,347,349,350,351,352,353,444,483,489,490,491,492,493,494,495,496,499,500,501,502,504,505,507,508,509,510,511,512,520,652,653,654,655,656,657,658,660,661,662,663,665,666,667,668,670,671,673,674,675,676,678,815,816,817,818,819,820,821,822,823,825,826,828,829,832,837,839,840,978,979,980,981,982,984,985,986,988,990,991,995,999,1001,1003,1004,1037,1038,1073,1075}
	datBase, err := models.NewDB()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	citsSlice, err := datBase.QueryCities(tsp150[:])
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	conns, err := datBase.QueryConnections(tsp150[:])
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
		Permutation:  tsp150[:],
		SolutionCost: 0.0, // Lo calculamos en el siguiente paso
		instance:     &instance,
	}
	
	// 1. Calculamos el costo base con la función oficial
	ts.SolutionCost = instance.Cost(&ts)

	// 2. Probamos 1000 intercambios aleatorios para estresar los casos borde
	for i := 0; i < 1000; i++ {
		u := instance.Random.IntN(150)
		v := instance.Random.IntN(150)
		for v == u {
			v = instance.Random.IntN(150)
		}

		// El vecino calcula el costo usando tu diferencial (balance)
		vecino := instance.Neighbour(&ts, u, v)

		// Verificamos calculando la ruta completa desde cero
		costoReal := instance.Cost(vecino)

		// Comparamos si el diferencial coincide con la realidad
		margen := math.Abs(vecino.Cost() - costoReal)
		if margen > 0.000001 {
			t.Fatalf("Fallo en swap(%d, %d): Delta dio %f, pero Cost dio %f", u, v, vecino.Cost(), costoReal)
		}
	}
}
