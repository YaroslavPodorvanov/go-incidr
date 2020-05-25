package incidr

import "sync/atomic"

type IPv4CidrList struct {
	data atomic.Value
}

func NewIPv4CidrList() *IPv4CidrList {
	var result = new(IPv4CidrList)

	result.update(nil)

	return result
}

func (l *IPv4CidrList) Contains(ip uint32) bool {
	var ranges = l.data.Load().([]IPv4Range)

	return IPv4Contains(ranges, ip)
}

func (l *IPv4CidrList) Update(ranges []IPv4Range) {
	l.update(ranges)
}

func (l *IPv4CidrList) update(ranges []IPv4Range) {
	l.data.Store(ranges)
}
