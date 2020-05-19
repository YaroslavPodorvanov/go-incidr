package incidr

import (
	"encoding/binary"
	"net"
)

type CidrList struct {
	ipv4 *IPv4CidrList
	ipv6 *IPv6CidrList
}

func NewCidrList(ipv4 *IPv4CidrList, ipv6 *IPv6CidrList) *CidrList {
	return &CidrList{ipv4: ipv4, ipv6: ipv6}
}

func (l *CidrList) Contains(ip net.IP) bool {
	{
		var ipv4 = ip.To4()

		if ipv4 != nil {
			return l.ipv4.Contains(binary.BigEndian.Uint32(ipv4))
		}
	}

	// just in case condition
	{
		var ipv6 = ip.To16()

		if ipv6 != nil {
			return l.ipv6.Contains(IPv6{
				binary.BigEndian.Uint64(ipv6[:8]),
				binary.BigEndian.Uint64(ipv6[8:]),
			})
		}
	}

	return false
}
