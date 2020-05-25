package incidr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIPv4Contains(t *testing.T) {
	require.Equal(t, false, IPv4Contains(nil, 1))

	var ranges = []IPv4Range{
		{
			First: 0,
			Last:  255,
		},
		{
			First: 1024,
			Last:  2047,
		},
	}

	require.Equal(t, true, IPv4Contains(ranges, 0))
	require.Equal(t, true, IPv4Contains(ranges, 1))
	require.Equal(t, true, IPv4Contains(ranges, 150))
	require.Equal(t, true, IPv4Contains(ranges, 255))

	require.Equal(t, false, IPv4Contains(ranges, 256))
	require.Equal(t, false, IPv4Contains(ranges, 1023))

	require.Equal(t, true, IPv4Contains(ranges, 1024))
	require.Equal(t, true, IPv4Contains(ranges, 1025))
	require.Equal(t, true, IPv4Contains(ranges, 1500))
	require.Equal(t, true, IPv4Contains(ranges, 2047))
}
