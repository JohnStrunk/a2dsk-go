// Package diskformats allows reading and writing sectors from emulated disk
// image formats.
package diskformats

// Block represents a block number on a disk
type Block int

// Sector represents a sector number on a track, typically [0, smax-1]
type Sector int

// Track represents a track number on a disk, typically [0, tmax-1]
type Track int

// BlockSize is the size of a block, in bytes
const BlockSize = 512

// BlockData is the contents read/written to a disk block
type BlockData []byte

// SectorSize is the size of a sector, in bytes
const SectorSize = 256

// SectorData is the contents read/written to a sector
type SectorData []byte

// Geometry represents the geometry of a disk
type Geometry struct {
	Sectors Sector
	Tracks  Track
}

// FiveQuarterDisk returns the geometry of a 5.25" disk
func FiveQuarterDisk() Geometry {
	return Geometry{Sectors: 16, Tracks: 35}
}

// SectorDisk is the common interface to the emulator formats that treat the
// underlying media as sectors (ala DOS3.3).
type SectorDisk interface {
	Read(s Sector, d []byte) error
	Write(s Sector, d []byte) error
}

// ProdosDisk is the common interface to emulator formats that can access data
// as ProDOS blocks
type ProdosDisk interface {
	Read(b Block, d []byte) error
	Write(b Block, d []byte) error
}

// SectorCount returns the total number of sectors in a Geometry
func (g Geometry) SectorCount() Sector {
	return g.Sectors * Sector(g.Tracks)
}

// Size returns the total number of bytes in a Geometry
func (g Geometry) Size() (bytes int) {
	return int(g.SectorCount()) * SectorSize
}
