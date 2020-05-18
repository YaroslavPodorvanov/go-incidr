package incidr

import "sort"

func In(ranges []Range, ip uint32) bool {
	length := len(ranges)
	index := sort.Search(length, func(i int) bool {
		return ip <= ranges[i].Last
	})

	if index < length {
		ispRange := ranges[index]

		return ispRange.First <= ip
	}

	return false
}
