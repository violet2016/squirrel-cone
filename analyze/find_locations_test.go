package analyze

import (
	"fmt"
	"testing"
)

func TestSearchForSig(t *testing.T) {
	searchForTest := SearchForSig(testRegion)

	index := searchForTest("ABABAB")
	if index != -1 {
		t.Fatal("It shouldn't found ABABAB")
	}
	index = searchForTest("72696E")
	if index <= 0 {
		fmt.Println(testRegion)
		t.Fatal("It should have found 72696E")
	}
}
