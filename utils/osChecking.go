package utils

import (
	"bytes"
	"os"
	"runtime"
)

func IsWSL() bool {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	return bytes.Contains(data, []byte("Microsoft")) || bytes.Contains(data, []byte("WSL"))
}

func IsWindows() bool {
	return runtime.GOOS == "windows" && !IsWSL()
}

func IsMac() bool {
	return runtime.GOOS == "darwin"
}

func IsLinux() bool {
	return runtime.GOOS == "linux" && !IsWSL()
}
