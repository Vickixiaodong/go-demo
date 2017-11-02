package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())
}

/*
Architecture: amd64
Go Version: go1.9
Operating System: darwin
GOPATH=/Users/xiexiaodong/Documents/vagrant/dev-tron/workspace
GOROOT=/usr/local/go
*/
