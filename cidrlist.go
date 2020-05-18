package incidr

import "sync/atomic"

type CidrList struct {
	data atomic.Value
}

func NewCidrList() *CidrList {
	var result = new(CidrList)

	result.update(nil)

	return result
}

func (l *CidrList) Contains(ip uint32) bool {
	var ranges = l.data.Load().([]Range)

	return In(ranges, ip)
}

func (l *CidrList) Update(ranges []Range) {
	l.update(ranges)
}

func (l *CidrList) update(ranges []Range) {
	l.data.Store(ranges)
}
