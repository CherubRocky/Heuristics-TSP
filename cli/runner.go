package cli

import(
	"github.com/CherubRocky/Heuristics-TSP/tsp_solver"
	"github.com/CherubRocky/Heuristics-TSP/models"
	"github.com/CherubRocky/Heuristics-TSP/tsp"

	"math/rand/v2"
	"slices"
	"fmt"
	"os"
)

func Run(seed uint64) {
	ids, err := readLine()
	ids2 := slices.Clone(ids)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	datBase, err := models.NewDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	citsSlice, err := datBase.QueryCities(ids)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	conns, err := datBase.QueryConnections(ids)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	// Close DB
	err = datBase.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	rng := rand.New(rand.NewPCG(seed, seed + 1))
	
	cities := tsp.NewCities(conns, citsSlice)
	instance := tsp_solver.Instance{cities, cities.KEdgesSum, rng}
	fmt.Println("Se llegó al punto")
	bestSol := instance.Solve(ids2)
	travelSol, _ := bestSol.(*tsp_solver.TravelSolution)
	fmt.Println("Mejor ruta encontrada:", travelSol.Permutation)
	fmt.Printf("Costo normalizado: %.6f\n", travelSol.Cost())
	

}
