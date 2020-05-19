package incidr

type IPv4Range struct {
	First uint32
	Last  uint32
}

type IPv4RangeSort []IPv4Range

func (c IPv4RangeSort) Len() int {
	return len(c)
}

func (c IPv4RangeSort) Less(i, j int) bool {
	a := c[i]
	b := c[j]

	if a.First < b.First {
		return true
	}

	if a.First > b.First {
		return false
	}

	// 0 to 127, before 0 to 15
	return a.Last > b.Last
}

func (c IPv4RangeSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
