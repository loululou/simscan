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
	ports := []int{21, 22, 25, 53, 80, 110, 443, 3306, 3389, 445, 7000, 7070, 8080, 8081, 8082, 8083, 8084, 8443, 10000, 9090}

	fmt.Printf("Scanning target: %s\n", target)
	internal.ScanRange(target, ports)
	
}