package cone

import (
	"fmt"
	"testing"
)

func TestLoadProcessRegions(t *testing.T) {
	h, _ := OpenProcess(pI.ProcessId)
	defer CloseHandle(h)
	regions := LoadProcessRegions(h)
	if len(regions) == 0 {
		t.Error("Region is empty")
	}
}
