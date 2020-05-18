package incidr

import (
	"encoding/binary"
	"math"
	"net"
)

func ParseCIDR(cidr string) (uint32, uint32, error) {
	_, ipNet, err := net.ParseCIDR(cidr)

	if err != nil {
		return 0, 0, err
	}

	ip := binary.BigEndian.Uint32(ipNet.IP)
	mask := binary.BigEndian.Uint32(ipNet.Mask)

	firstIP := ip & mask

	lastIP := firstIP | (math.MaxUint32 ^ mask)

	return firstIP, lastIP, nil
}
