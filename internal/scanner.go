package internal

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Scan a range of IPs
func ScanRange(ipRange string, ports []int) {
	ips, err := ParseIPRange(ipRange)
	if err != nil {
		fmt.Println("[-] Invalid IP range:", err)
		return
	}

	var wg sync.WaitGroup
	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			scanHost(ip, ports)
		}(ip)
	}
	wg.Wait()
}

func scanHost(ip string, ports []int) {
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			fmt.Printf("[+] %s:%d is open\n", ip, port)
			conn.Close()
		}
	}
}

// Parse IP range (CIDR or explicit range)
func ParseIPRange(ipRange string) ([]string, error) {
	var ips []string

	// Check if input is in CIDR format
	if strings.Contains(ipRange, "/") {
		ip, ipnet, err := net.ParseCIDR(ipRange)
		if err != nil {
			return nil, err
		}

		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			ips = append(ips, ip.String())
		}
	} else if strings.Contains(ipRange, "-") {
		// Handle range input (e.g., 192.168.1.1-192.168.1.100)
		parts := strings.Split(ipRange, "-")
		startIP := net.ParseIP(parts[0])
		endIP := net.ParseIP(parts[1])

		if startIP == nil || endIP == nil {
			return nil, fmt.Errorf("invalid IP address range")
		}

		for ip := startIP; !ip.Equal(endIP); inc(ip) {
			ips = append(ips, ip.String())
		}
		ips = append(ips, endIP.String()) // Add the last IP
	} else {
		ips = append(ips, ipRange) // Single IP input
	}

	return ips, nil
}

// Increment IP address
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
