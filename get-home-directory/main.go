package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func main() {
	// Home directory: /Users/xiexiaodong
	fmt.Println("Home directory:", homeDir())

	// OS: darwin
	fmt.Println("OS:", runtime.GOOS)
}

// Get home directory
func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}

	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}

	return ""
}
