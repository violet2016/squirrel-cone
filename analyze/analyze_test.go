package analyze
import (
	"testing"
	"os"
)

var testRegion []byte

func initTestRegion() {
	testRegion = []byte("I'm a test string. I don't have too much to say.")
}
func TestMain(m *testing.M) {
	initTestRegion()
	retcode := m.Run()
	os.Exit(retcode)
}