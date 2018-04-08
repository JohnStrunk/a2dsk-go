package diskformats

import (
	"errors"
	"io"
)

// DoDisk is a DOS 3.3 ordered disk.
type DoDisk struct {
	file io.ReadWriteSeeker
}

// DoDisk isa SectorDisk
var _ SectorDisk = (*DoDisk)(nil)

// NewDoDisk returns a DOS-ordered disk on top of the provided I/O stream
func NewDoDisk(file io.ReadWriteSeeker) *DoDisk {
	return &DoDisk{file: file}
}

// Read returns the data from a given disk sectorData
func (d *DoDisk) Read(s Sector, data []byte) error {
	if len(data) < SectorSize {
		return errors.New("read: buffer too short")
	}
	if _, err := d.file.Seek(int64(s)*SectorSize, io.SeekStart); err != nil {
		return err
	}
	b, err := d.file.Read(data[:SectorSize])
	if b != SectorSize {
		if err != nil {
			return err
		}
		return errors.New("read: unable to read sector, unknown error")
	}
	return nil
}

// Write stores the provided data into the specified disk sector
func (d *DoDisk) Write(s Sector, data []byte) error {
	if len(data) < SectorSize {
		return errors.New("write: buffer too short")
	}
	return nil
}
