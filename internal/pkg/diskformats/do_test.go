package diskformats

import (
	"errors"
	"testing"
)

type FakeRWS struct {
	Bytes int
	Err   error
}

func (f FakeRWS) Read([]byte) (n int, e error) {
	return f.Bytes, f.Err
}
func (f FakeRWS) Write([]byte) (n int, e error) {
	return f.Bytes, f.Err
}
func (f FakeRWS) Seek(o int64, w int) (n int64, e error) {
	return int64(f.Bytes), f.Err
}
func TestDoDiskRead(t *testing.T) {
	r := FakeRWS{Bytes: SectorSize, Err: nil}
	d := NewDoDisk(r)
	data := make([]byte, SectorSize)
	rv := d.Read(0, data)
	if rv != nil {
		t.Errorf("read should have returned nil")
	}
}

func TestDoDiskReadErrors(t *testing.T) {
	r := FakeRWS{Bytes: SectorSize, Err: nil}
	d := NewDoDisk(r)
	rv := d.Read(0, nil)
	if rv == nil {
		t.Errorf("DoDisk.Read: nil buffer should return error")
	}

	data := make([]byte, 10)
	rv = d.Read(0, data)
	if rv == nil {
		t.Errorf("DoDisk.Read: short buffer should return error")
	}

	data = make([]byte, SectorSize)
	r.Err = errors.New("seek error")
	d = NewDoDisk(r)
	rv = d.Read(3, data)
	if rv == nil {
		t.Errorf("DoDisk.Read: failed seek should return error")
	}

	r.Err = nil
	r.Bytes = 5 // Short read
	d = NewDoDisk(r)
	rv = d.Read(3, data)
	if rv == nil {
		t.Errorf("DoDisk.Read: failed short read should return error")
	}
}

func TestDoDiskWriteInvalidBuffer(t *testing.T) {
	r := FakeRWS{Bytes: SectorSize, Err: nil}
	d := NewDoDisk(r)
	rv := d.Write(0, nil)
	if rv == nil {
		t.Errorf("DoDisk.Write: nil buffer should return error")
	}

	data := make([]byte, 10)
	rv = d.Write(0, data)
	if rv == nil {
		t.Errorf("DoDisk.Write: short buffer should return error")
	}
}
