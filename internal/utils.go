// This is for validation
package internal

import (
	"net"
)

// Validation
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}