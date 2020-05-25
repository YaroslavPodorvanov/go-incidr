package incidr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIPv6Contains(t *testing.T) {
	require.Equal(t, false, IPv6Contains(nil, IPv6{0, 0}))

	var ranges = []IPv6Range{
		{
			First: IPv6{0, 0},
			Last:  IPv6{0, 255},
		},

		{
			First: IPv6{1024, 0},
			Last:  IPv6{1024, 255},
		},
	}

	require.Equal(t, true, IPv6Contains(ranges, IPv6{0, 0}))
	require.Equal(t, true, IPv6Contains(ranges, IPv6{0, 1}))
	require.Equal(t, true, IPv6Contains(ranges, IPv6{0, 255}))

	require.Equal(t, false, IPv6Contains(ranges, IPv6{0, 256}))
	require.Equal(t, false, IPv6Contains(ranges, IPv6{1, 0}))

	require.Equal(t, true, IPv6Contains(ranges, IPv6{1024, 0}))
	require.Equal(t, true, IPv6Contains(ranges, IPv6{1024, 1}))
	require.Equal(t, true, IPv6Contains(ranges, IPv6{1024, 255}))
}
