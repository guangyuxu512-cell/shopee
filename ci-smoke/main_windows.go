package main

import (
	"os"
	"syscall"
	"unsafe"
)

// This is intentionally a harmless, dependency-free CI/CD delivery smoke test.
func main() {
	if len(os.Args) == 3 && os.Args[1] == "--self-test" {
		if err := os.WriteFile(os.Args[2], []byte("CI_SMOKE_OK\n"), 0600); err != nil {
			os.Exit(1)
		}
		return
	}

	title, err := syscall.UTF16PtrFromString("ChatGPT CI/CD delivery test")
	if err != nil {
		os.Exit(1)
	}
	message, err := syscall.UTF16PtrFromString("Success! This Windows EXE was built by GitHub Actions and downloaded as a ZIP.")
	if err != nil {
		os.Exit(1)
	}

	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")
	result, _, _ := messageBox.Call(
		0,
		uintptr(unsafe.Pointer(message)),
		uintptr(unsafe.Pointer(title)),
		0x40,
	)
	if result == 0 {
		os.Exit(1)
	}
}
