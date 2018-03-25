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
