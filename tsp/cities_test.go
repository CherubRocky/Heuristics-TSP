package tsp
import (
	"testing"
	"math"
)

func TestNaturalDistance(t *testing.T) {
	cdmx := NewCoordinates(19.4342, -99.1386)
	tampico := NewCoordinates(22.2167, -97.85)

	got := getNaturalDistance(cdmx, tampico)
	want := 337239.25
	margin := math.Abs(got - want)
	if margin > 0.000001 {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestToRadians(t *testing.T) {
	got := toRadians(360)
	want := 2 * math.Pi
	margin := math.Abs(got - want)
	if margin > 0.000001 {
		t.Errorf("got %f, want %f", got, want)
	}
}
