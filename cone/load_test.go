package cone

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadProcessRegions(t *testing.T) {
	regions := LoadProcessRegions(globalHandle)
	if len(regions) == 0 {
		t.Fatal("Region is empty")
	}
	files := SaveRegions(globalHandle, regions)(os.TempDir())
	if len(files) == 0 {
		t.Fatal("No files saved")
	}
	buf, err := ioutil.ReadFile(files[0])
	FatalOnErr(t, err)
	if len(buf) == 0 {
		t.Fatal("read file is empty")
	}
	for _, f := range files {
		os.Remove(f)
	}
}
