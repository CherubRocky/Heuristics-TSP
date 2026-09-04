package tsp

type Connection struct {
	Cit1 City
	Cit2 City
	Distance float64
}

type City struct {
	Id int
	Coords Coordinates
}

// Returns first the max and then the min id's
func (c Connection) getMaxAndMin() (int, int) {
	return getMaxAndMin(c.Cit1.Id, c.Cit2.Id)
}

func getMaxAndMin(one, two int) (int, int) {
	if one > two {
		return one, two
	}
	return two, one
}
