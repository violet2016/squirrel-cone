package main

import (
	"flag"
	"os"
	"path/filepath"
	"squirrel-cone/cone"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	pidPtr := flag.Int("pid", 0, "the pid of ffxiv")
	flag.Parse()
	pid := *pidPtr
	if pid != 0 {
		saveAllRegions(uint32(pid))
	}
}
func saveAllRegions(pid uint32) {
	h, err := cone.OpenProcess(pid)
	defer cone.CloseHandle(h)
	PanicOnErr(err)
	regions := cone.LoadProcessRegions(h)
	if len(regions) == 0 {
		panic("no regions found")
	}
	prefix := filepath.Join(os.TempDir(), "ffxiv")
	os.Mkdir(prefix, 0644)
	cone.SaveRegions(h, regions)(prefix)
}
