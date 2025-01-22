// This is the entry point

package main

import (
	"fmt"
	"os"
	"github.com/loululou/simscan/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: simscan <target>")
		os.Exit(1)
	}

	target := os.Args[1]
	fmt.Println("Scanning target: %s\n", target)
	internal.ScanTarget(target)
}