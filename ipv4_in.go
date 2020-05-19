package incidr

import "sort"

func IPv4In(ranges []IPv4Range, ip uint32) bool {
	length := len(ranges)
	index := sort.Search(length, func(i int) bool {
		return ip <= ranges[i].Last
	})

	if index < length {
		ipRange := ranges[index]

		return ipRange.First <= ip
	}

	return false
}
