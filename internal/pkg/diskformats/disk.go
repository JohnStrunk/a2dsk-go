// Package diskformats allows reading and writing sectors from emulated disk
// image formats.
package diskformats

// Track represents a track number on a disk, typically [0, tmax-1]
type Track int

// Sector represents a sector number on a track, typically [0, smax-1]
type Sector int

// Size in bytes of a sector
const SectorSize = 256

// SectorData is the contents read/written to a sector
type SectorData []byte

// Geometry represents the geometry of a disk
type Geometry struct {
	Sectors Sector
	Tracks  Track
}

// Geometry of a 5.25" disk
func FiveQuarterDisk() Geometry {
	return Geometry{Sectors: 16, Tracks: 35}
}

// SectoredDisk is the common interface to the emulator formats that treat the
// underlying media as cooked sectors.
type SectoredDisk interface {
	read(t Track, s Sector, d SectorData) error
	write(t Track, s Sector, d SectorData) error
	close() error
}

func (g Geometry) SectorCount() Sector {
	return g.Sectors * Sector(g.Tracks)
}

func (g Geometry) Size() (bytes int) {
	return int(g.SectorCount()) * SectorSize
}
