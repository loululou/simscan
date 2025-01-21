package internal

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ScanTarget scans a single IP or an IP range in CIDR notation.
func ScanTarget(target string) {
	ips, err := getIPAddresses(target)
	if err != nil {
		fmt.Printf("[-] Invalid IP address or range: %s\n", target)
		return
	}

	ports := []int{21, 22, 25, 53, 80, 110, 443, 3306, 3389}
	var wg sync.WaitGroup

	for _, ip := range ips {
		for _, port := range ports {
			wg.Add(1)
			go func(ip string, port int) {
				defer wg.Done()
				address := fmt.Sprintf("%s:%d", ip, port)
				conn, err := net.DialTimeout("tcp", address, 2*time.Second)
				if err == nil {
					fmt.Printf("[+] Port %d is open on %s\n", port, ip)
					conn.Close()
				} else {
					fmt.Printf("[-] Port %d is closed on %s\n", port, ip)
				}
			}(ip, port)
		}
	}

	wg.Wait()
}

// getIPAddresses handles single IPs and CIDR ranges, returning a list of IP addresses to scan.
func getIPAddresses(target string) ([]string, error) {
	var ips []string
	ip, ipNet, err := net.ParseCIDR(target)
	if err != nil {
		// Check if it's a valid single IP
		if net.ParseIP(target) == nil {
			return nil, fmt.Errorf("invalid IP address or range")
		}
		return []string{target}, nil
	}

	// Iterate through all IPs in the CIDR range
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// Remove network and broadcast addresses for non-/32 subnets
	if len(ips) > 2 {
		return ips[1 : len(ips)-1], nil
	}

	return ips, nil
}

// inc increments an IP address to move to the next one in range
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

