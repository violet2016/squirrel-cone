package cone

import (
	"os"
	"syscall"
	"testing"
)

var sI syscall.StartupInfo
var pI syscall.ProcessInformation

func openCalc() {
	argv := syscall.StringToUTF16Ptr("c:\\windows\\system32\\calc.exe")
	syscall.CreateProcess(
		nil,
		argv,
		nil,
		nil,
		true,
		0,
		nil,
		nil,
		&sI,
		&pI)
}
func terminateCalc() {
	syscall.TerminateProcess(pI.Process, 0)
}
func TestMain(m *testing.M) {
	openCalc()
	retCode := m.Run()
	terminateCalc()
	os.Exit(retCode)
}
