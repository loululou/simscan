package main

import (
	"fmt"
	"os"
	"simscan/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: simscan <target>")
		os.Exit(1)
	}

	target := os.Args[1]
	fmt.Printf("Scanning target: %s\n", target)
	internal.ScanTarget(target)
}

