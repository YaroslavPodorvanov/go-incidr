package incidr

import "sort"

func RangeSequenceByIPv6CIDRs(cidrs []string) []IPv6Range {
	length := len(cidrs)

	if length == 0 {
		return nil
	}

	all := make([]IPv6Range, 0, length)

	for _, cidr := range cidrs {
		first, last, err := ParseIPv6CIDR(cidr)

		if err != nil {
			// NOP

			continue
		}

		all = append(all, IPv6Range{
			First: first,
			Last:  last,
		})
	}

	// O(N*ln(N))
	sort.Sort(IPv6RangeSort(all))

	unique := all[:1]

	current := all[0]

	// O(N)
	for i := 1; i < length; i++ {
		next := all[i]

		if IPv6Less(current.Last, next.Last) {
			unique = append(unique, next)

			current = next
		}
	}

	return unique
}
