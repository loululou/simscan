// Basic function
package internal

import (
	"net"
	"fmt"
	"time"
)

func BasicScanner(target string){
	ports := []int{21, 22, 25, 53, 80, 110, 443, 3306, 3389, 445, 8080, 8081, 8082, 8083, 8084, 10000}
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err == nil {
			fmt.Printf("--- Scanning %s ---)
			fmt.Printf("[+] Port %d is open on %s\n", port, target)
			conn.Close()
		} else {
			fmt.Printf("[-] Port %d is closed on %s\n", port, target)
		}
	}
}

