package incidr

import "sort"

func IPv6Contains(ranges []IPv6Range, ip IPv6) bool {
	length := len(ranges)
	index := sort.Search(length, func(i int) bool {
		return IPv6LessEqual(ip, ranges[i].Last)
	})

	if index < length {
		ipRange := ranges[index]

		return IPv6LessEqual(ipRange.First, ip)
	}

	return false
}
