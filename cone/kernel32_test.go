package cone

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unsafe"
)

func FatalOnErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal("Error occurred", err)
	}
}
func TestVirtualQueryEx(t *testing.T) {
	info := MemoryBasicInformation{}
	buf, bufSize := VirtualQueryEx(globalHandle, 0, unsafe.Sizeof(info))
	if bufSize != unsafe.Sizeof(info) {
		t.Fatal("buf size wrong")
	}
	err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &info)
	FatalOnErr(t, err)
	if info.RegionSize == 0 {
		t.Fatal("RegionSize is 0")
	}
}
