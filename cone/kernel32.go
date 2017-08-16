// Read Windows Process Memory
package cone

import (
	"syscall"
	"unsafe"
)

// required functions in kernel32.dll
var (
	modkernel32           = syscall.NewLazyDLL("kernel32.dll")
	procReadProcessMemory = modkernel32.NewProc("ReadProcessMemory")
	procVirtualQueryEx    = modkernel32.NewProc("VirtualQueryEx")
)

const PROCESS_VM_ALL = 0x001F0FFF

type MemoryBasicInformation64 struct {
	BaseAddress       uint64
	AllocationBase    uint64
	AllocationProtect uint32
	Alignment         uint32
	RegionSize        uint64
	State             uint32
	Protect           uint32
	Type              uint32
	Alignment2        uint32
}

type MemoryBasicInformation32 struct {
	BaseAddress       uint32
	AllocationBase    uint32
	AllocationProtect uint32
	RegionSize        uint32
	State             uint32
	Protect           uint32
	Type              uint32
}

type MemoryBasicInformation MemoryBasicInformation64

func OpenProcess(pid uint32) (syscall.Handle, error) {
	return syscall.OpenProcess(PROCESS_VM_ALL, false, pid)
}

func CloseHandle(handle syscall.Handle) error {
	return syscall.CloseHandle(handle)
}

func ReadProcessMemory(hProcess syscall.Handle, lpBaseAddress, nSize uintptr) (lpBuffer []byte, lpNumberOfBytesRead int, ok bool) {

	var nBytesRead int
	buf := make([]byte, nSize)
	ret, _, _ := procReadProcessMemory.Call(
		uintptr(hProcess),
		lpBaseAddress,
		uintptr(unsafe.Pointer(&buf[0])),
		nSize,
		uintptr(unsafe.Pointer(&nBytesRead)),
	)
	return buf, nBytesRead, ret != 0
}

func VirtualQueryEx(hProcess syscall.Handle, lpAddress, nSize uintptr) ([]byte, uintptr) {
	infoBuf := make([]byte, nSize)
	nSizeOfInfoBuffer, _, _ := procVirtualQueryEx.Call(
		uintptr(hProcess),
		lpAddress,
		uintptr(unsafe.Pointer(&infoBuf[0])),
		nSize,
	)
	return infoBuf, nSizeOfInfoBuffer
}
