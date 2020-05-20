package incidr

type IPv6 [2]uint64

type IPv6Range struct {
	First IPv6
	Last  IPv6
}

type IPv6RangeSort []IPv6Range

func (c IPv6RangeSort) Len() int {
	return len(c)
}

func (c IPv6RangeSort) Less(i, j int) bool {
	a := c[i]
	b := c[j]

	if IPv6Less(a.First, b.First) {
		return true
	}

	if IPv6Less(b.First, a.First) {
		return false
	}

	// 0 to 127, before 0 to 15
	return IPv6Less(b.Last, a.Last)
}

func (c IPv6RangeSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func IPv6Less(a, b IPv6) bool {
	if a[0] < b[0] {
		return true
	}

	if a[0] > b[0] {
		return false
	}

	return a[1] < b[1]
}

func IPv6LessEqual(a, b IPv6) bool {
	if a[0] < b[0] {
		return true
	}

	if a[0] > b[0] {
		return false
	}

	return a[1] <= b[1]
}
