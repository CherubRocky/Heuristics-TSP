package main

import (
	"github.com/CherubRocky/Heuristics-TSP/cli"
	"flag"
)

func main() {
	seedFlag := flag.Uint64("seed", 4, "Semilla para el generador pseudoaleatorio")
	flag.Parse()
	seed := *seedFlag
	cli.Run(seed)
}
