package pkg

import (
    "net"
)

func IsValidIP(ip string) bool {
    return net.ParseIP(ip) != nil
}
