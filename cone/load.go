package cone

import (
	"bytes"
	"encoding/binary"
	"syscall"
	"unsafe"
)

const MemCommit = 0x1000
const PageGuard = 0x100
const Writable = syscall.PAGE_READWRITE | syscall.PAGE_WRITECOPY | syscall.PAGE_EXECUTE_READWRITE | syscall.PAGE_EXECUTE_WRITECOPY | PageGuard

func LoadProcessRegions(handle syscall.Handle) []MemoryBasicInformation {
	address := uintptr(0)
	regions := []MemoryBasicInformation{}
	for {
		info := MemoryBasicInformation{}
		buf, bufSize := VirtualQueryEx(handle, address, unsafe.Sizeof(info))
		if bufSize == 0 {
			break
		}
		binary.Read(bytes.NewReader(buf), binary.LittleEndian, &info)
		if info.State&MemCommit != 0 && info.Protect&Writable != 0 && info.Protect&PageGuard == 0 {
			regions = append(regions, info)
		}
		address = uintptr(info.BaseAddress + info.RegionSize)
	}

	return regions
}
