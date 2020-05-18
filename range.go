package incidr

type Range struct {
	First uint32
	Last  uint32
}

type RangeSort []Range

func (c RangeSort) Len() int {
	return len(c)
}

func (c RangeSort) Less(i, j int) bool {
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

func (c RangeSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
