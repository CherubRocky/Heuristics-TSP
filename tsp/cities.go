package tsp

import (
	"math"
)

type Cities struct {
	Matrix [][]float64
	KEdgesSum float64
}

// We can implement an inferior triangular matrix with a one-dimensional array
func NewCities(maxId int) *Cities {
	cities := Cities{Matrix: make([][]float64, maxId + 1)}
	for i := range cities.Matrix {
		cities.Matrix[i] = make([]float64, maxId + 1)
	}
	return &cities
}

func (cities *Cities) GetDistance(id1 int, id2 int) float64 {
	return cities.Matrix[id1][id2]
}

func getNaturalDistance(coord1 coordinates, coord2 coordinates) float64 {
	b := math.Pow(math.Sin((coord1.latitude - coord2.latitude) / 2), 2)
	c := math.Pow(math.Sin((coord1.longitude - coord2.longitude) / 2), 2)
	a := b + math.Cos(coord1.latitude) * math.Cos(coord2.latitude) * c
	d := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(math.Abs(1 - a)))
	return 6373000 * d
}

type coordinates struct {
	latitude float64
	longitude float64
}

// One must have in mind that in the db, id_city_2 > id_city_1 in the connections table
func newCoordinates(degLat float64, degLon float64) coordinates {
	latitude := toRadians(degLat)
	longitude := toRadians(degLon)
	return coordinates{latitude, longitude}
}

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
