package tsp

import (	
	"cmp"
	"math"
	"errors"
	"slices"
)

type Cities struct {
	Matrix [][]float64
	KEdgesSum float64
}

// It is always id_city_1 > id_city_j (superior triangular matrix)
// We can implement an inferior triangular matrix with a one-dimensional array
func NewCities(conns []Connection, cits []City) *Cities {
	maxId, _ := getMax(cits)
	cities := Cities{Matrix: make([][]float64, maxId + 1)}
	for i := range cities.Matrix {
		cities.Matrix[i] = make([]float64, maxId + 1)
	}
	cities.fillBasic(conns)
	cities.calcKEdgesSum(conns, len(cits))
	cities.fillPhantom(cits)
	return &cities
}

// Assuming the lightest connections are always the ones in the DB
func (c *Cities) calcKEdgesSum (conns []Connection, citNum int) {
	slices.SortFunc(conns, func(c1, c2 Connection) int {
		return cmp.Compare(c1.Distance, c2.Distance)
	})
	var sum float64
	for i := len(conns) - 1; i > len(conns) - citNum && i >= 0; i-- {
		sum += conns[i].Distance
	}
	c.KEdgesSum = sum
}

func (c *Cities) fillPhantom(citsArr []City) {
	for i := 0; i < len(citsArr); i++ {
		for j := i + 1; j < len(citsArr); j++ {
			// We know there is no row in Connections with distance field equal to zero
			if c.GetDistance(citsArr[i].Id, citsArr[j].Id) == 0 {
				natDist := getNaturalDistance(citsArr[i].Coords, citsArr[j].Coords)
				c.SetDistance(citsArr[i].Id, citsArr[j].Id, natDist)
			}
		}
	}
}

func (c *Cities) fillBasic(connections []Connection) {
	for _, conn := range connections {
		c.SetDistance(conn.Cit1.Id, conn.Cit2.Id, conn.Distance)
	}
} 

func (cities *Cities) GetDistance(id1 int, id2 int) float64 {
	big, lit := getMaxAndMin(id1, id2)
	return cities.Matrix[lit][big]
}

func (cities *Cities) SetDistance(id1 int, id2 int, dist float64) {
	big, lit := getMaxAndMin(id1, id2)
	cities.Matrix[lit][big] = dist
}

func getNaturalDistance(coord1 Coordinates, coord2 Coordinates) float64 {
	b := math.Pow(math.Sin((coord1.Latitude - coord2.Latitude) / 2), 2)
	c := math.Pow(math.Sin((coord1.Longitude - coord2.Longitude) / 2), 2)
	a := b + math.Cos(coord1.Latitude) * math.Cos(coord2.Latitude) * c
	d := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(math.Abs(1 - a)))
	return 6373000 * d
}

type Coordinates struct {
	Latitude float64
	Longitude float64
}

// One must have in mind that in the db, id_city_2 > id_city_1 in the connections table
func NewCoordinates(degLat float64, degLon float64) Coordinates {
	latitude := toRadians(degLat)
	longitude := toRadians(degLon)
	return Coordinates{latitude, longitude}
}

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func getMax(arr []City) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("Array is length 0. No maximum here.")
	}
	max := arr[0].Id
	for _, c := range arr {
		if max < c.Id {
			max = c.Id
		}
	}
	return max, nil
}
