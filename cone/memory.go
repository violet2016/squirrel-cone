package cone

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"syscall"
)

const MemFileMode = 0644

func SaveRegions(h syscall.Handle, regions []MemoryBasicInformation) func(string) ([]string, []MemoryBasicInformation) {
	return func(prefix string) ([]string, []MemoryBasicInformation) {
		files := []string{}
		failed := []MemoryBasicInformation{}
		for _, r := range regions {
			ch := make(chan string)
			go ReadRegionAndSaveToFile(h, r, prefix, ch)
			f := <-ch
			if len(f) > 0 {
				files = append(files, f)
			} else {
				failed = append(failed, r)
			}
		}
		return files, failed
	}
}

func ReadRegionAndSaveToFile(h syscall.Handle, r MemoryBasicInformation, prefix string, ch chan string) {
	buf, _, ret := ReadProcessMemory(h, uintptr(r.BaseAddress), uintptr(r.RegionSize))
	filename := filepath.Join(prefix, fmt.Sprintf("%x", r.BaseAddress))
	if !ret {
		log.Println("Failed to read memory for", filename, "size is:", r.RegionSize)
		ch <- ""
	} else {
		err := ioutil.WriteFile(filename, buf, MemFileMode)
		if err != nil {
			log.Panic("Cannot write file", err)
		}
		ch <- filename
	}
}
