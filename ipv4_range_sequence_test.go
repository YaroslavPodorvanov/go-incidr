package incidr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRangeSequenceByCIDRs(t *testing.T) {
	require.Equal(t, []IPv4Range(nil), RangeSequenceByIPv4CIDRs(nil))

	require.Equal(
		t,
		[]IPv4Range{
			{
				First: 0x0,
				Last:  0xffffffff,
			},
		},
		RangeSequenceByIPv4CIDRs([]string{"0.0.0.0/0"}),
	)

	require.Equal(
		t,
		[]IPv4Range{
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
		RangeSequenceByIPv4CIDRs([]string{
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
