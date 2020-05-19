package incidr

import (
	"sync/atomic"
)

type IPv6CidrList struct {
	data atomic.Value
}

func NewIPv6CidrList() *IPv6CidrList {
	var result = new(IPv6CidrList)

	result.update(nil)

	return result
}

func (l *IPv6CidrList) Contains(ip IPv6) bool {
	var ranges = l.data.Load().([]IPv6Range)

	return IPv6In(ranges, ip)
}

func (l *IPv6CidrList) Update(ranges []IPv6Range) {
	l.update(ranges)
}

func (l *IPv6CidrList) update(ranges []IPv6Range) {
	l.data.Store(ranges)
}
