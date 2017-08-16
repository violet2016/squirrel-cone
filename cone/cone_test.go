package cone

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

var sI syscall.StartupInfo
var pI syscall.ProcessInformation
var globalHandle syscall.Handle

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
	globalHandle, _ = OpenProcess(pI.ProcessId)
}

func terminateCalc() {
	err := CloseHandle(globalHandle)
	if err != nil {
		fmt.Println("Failed to close handle at the end of test", err)
	}
	err = syscall.TerminateProcess(pI.Process, 0)
	if err != nil {
		fmt.Println("Failed to terminate Calc at the end of test", err)
	}
}
func TestMain(m *testing.M) {
	openCalc()
	retCode := m.Run()
	terminateCalc()
	os.Exit(retCode)
}
