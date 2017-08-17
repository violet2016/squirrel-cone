package analyze

import "testing"

func TestLoadSigs(t *testing.T) {
	sigs := LoadSigsFile("../sigs/latest.json")
	if len(sigs) == 0 {
		t.Fatal("Load sigs is empty")
	}
}
