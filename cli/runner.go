package cli

import(
	"github.com/CherubRocky/Heuristics-TSP/tsp_solver"
	"github.com/CherubRocky/Heuristics-TSP/models"
	"github.com/CherubRocky/Heuristics-TSP/tsp"
	"slices"
	"fmt"
	"os"
)

func Run() {
	ids, err := readArguments()
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
	
	cities := tsp.NewCities(conns, citsSlice)
	instance := tsp_solver.Instance{cities, cities.KEdgesSum}
	sol := tsp_solver.TravelSolution{ids2, 0.0}
	distance := instance.Cost(&sol)
	fmt.Sprint("The distance of the given solution is: ")
	fmt.Sprintln(distance)
}
