package incidr

import "sort"

func RangeSequenceByCIDRs(cidrs []string) []Range {
	length := len(cidrs)

	if length == 0 {
		return nil
	}

	all := make([]Range, 0, length)

	for _, cidr := range cidrs {
		first, last, err := ParseCIDR(cidr)

		if err != nil {
			// NOP

			continue
		}

		all = append(all, Range{
			First: first,
			Last:  last,
		})
	}

	// O(N*ln(N))
	sort.Sort(RangeSort(all))

	unique := all[:1]

	current := all[0]

	// O(N)
	for i := 1; i < length; i++ {
		next := all[i]

		if next.Last > current.Last {
			unique = append(unique, next)

			current = next
		}
	}

	return unique
}
