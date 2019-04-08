// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package hcsshim

import (
	"syscall"
	"unsafe"

	"github.com/Microsoft/hcsshim/internal/interop"
	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modiphlpapi = windows.NewLazySystemDLL("iphlpapi.dll")

	procSetCurrentThreadCompartmentId = modiphlpapi.NewProc("SetCurrentThreadCompartmentId")
)

func SetCurrentThreadCompartmentId(compartmentId uint32) (hr error) {
	r0, _, _ := syscall.Syscall(procSetCurrentThreadCompartmentId.Addr(), 1, uintptr(compartmentId), 0, 0)
	if int32(r0) < 0 {
		hr = interop.Win32FromHresult(r0)
	}
	return
}