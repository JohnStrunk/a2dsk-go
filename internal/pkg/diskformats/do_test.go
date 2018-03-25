package diskformats

import (
	"testing"
)

func TestDoDiskRead(t *testing.T) {
	d := new(DoDisk)
	data := make([]byte, SectorSize)
	rv := d.read(0, 0, data)
	if rv != nil {
		t.Errorf("read should have returned nil")
	}
}

func TestDoDiskReadInvalidBuffer(t *testing.T) {
	d := new(DoDisk)
	rv := d.read(0, 0, nil)
	if rv == nil {
		t.Errorf("DoDisk.Read: nil buffer should return error")
	}

	data := make([]byte, 10)
	rv = d.read(0, 0, data)
	if rv == nil {
		t.Errorf("DoDisk.Read: short buffer should return error")
	}
}
