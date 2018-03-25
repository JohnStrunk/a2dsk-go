package diskformats

import (
	"testing"
)

var sectorData = []struct {
	t  Track
	s  Sector
	sc Sector
	b  int
}{
	{0, 0, 0, 0},
	{1, 1, 1, 256},
	{35, 16, 560, 143360}, // 5.25" disk
}

func TestGeometrySectorCount(t *testing.T) {
	for _, d := range sectorData {
		g := Geometry{Tracks: d.t, Sectors: d.s}
		sc := g.SectorCount()
		if sc != d.sc {
			t.Errorf("SectorCount of %v/%v: want %v, got %v",
				d.t, d.s, d.sc, sc)
		}
	}
}

func TestGeometrySize(t *testing.T) {
	for _, d := range sectorData {
		g := Geometry{Tracks: d.t, Sectors: d.s}
		b := g.Size()
		if b != d.b {
			t.Errorf("SectorCount of %v/%v: want %v, got %v",
				d.t, d.s, d.b, b)
		}
	}
}
