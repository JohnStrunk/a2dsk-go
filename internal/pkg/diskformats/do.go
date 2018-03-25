package diskformats

import (
	"errors"
	"io"
)

// DoDisk is a DOS 3.3 ordered disk.
type DoDisk struct {
	file io.ReadWriteCloser
}

func (d *DoDisk) read(t Track, s Sector, data SectorData) error {
	if cap(data) < SectorSize {
		return errors.New("read: invalid buffer")
	}
	return nil
}

//18 ‣‧‧‧‧‧‧‧write(t Track, s Sector, d SectorData) error
//19 ‣‧‧‧‧‧‧‧close() error
