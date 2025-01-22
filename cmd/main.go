// This is the entry point

package main

import (
	"fmt"
	"os"
	"github.com/loululou/simscan/internal"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: simscan <target>")
		os.Exit(1)
	}

	target := os.Args[1]
	ports := []int{21, 22, 25, 53, 80, 110, 443, 3306, 3389, 445, 8080, 8081, 8082, 8083, 8084, 10000, 9090}

	fmt.Println("Scanning target:", target)
	internal.ScanRange(target, ports)
}