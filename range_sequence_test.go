package incidr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRangeSequenceByCIDRs(t *testing.T) {
	require.Equal(t, []Range(nil), RangeSequenceByCIDRs(nil))

	require.Equal(
		t,
		[]Range{
			{
				First: 0x0,
				Last:  0xffffffff,
			},
		},
		RangeSequenceByCIDRs([]string{"0.0.0.0/0"}),
	)

	require.Equal(
		t,
		[]Range{
			{
				First: 0x01010100,
				Last:  0x010101ff,
			},
			{
				First: 0x01010200,
				Last:  0x010102ff,
			},
			{
				First: 0x01010300,
				Last:  0x0101037f,
			},
		},
		RangeSequenceByCIDRs([]string{
			// 2
			"1.1.2.0/25", // will merge
			"1.1.2.0/24", // will merge

			// 1
			"1.1.1.0/24", // will merge
			"1.1.1.0/25", // will merge
			"1.1.1.0/26", // will merge
			"1.1.1.0/27", // will merge
			"1.1.1.0/28", // will merge

			// 3
			"1.1.3.0/25",
		}),
	)
}
