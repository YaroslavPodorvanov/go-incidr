package incidr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	assertParseCIDR(t, "104.131.0.0/16", 0x68830000, 0x6883ffff)
	assertParseCIDR(t, "104.219.251.162/32", 0x68dbfba2, 0x68dbfba2)
	assertParseCIDR(t, "1.2.64.0/18", 0x1024000, 0x1027fff)
	assertParseCIDR(t, "0.0.0.0/0", 0x0, 0xffffffff)
}

func TestParseInvalidCIDRAddress(t *testing.T) {
	assertExpectErrParseCIDR(t, "104.131.0.0/", "invalid CIDR address: 104.131.0.0/")
	assertExpectErrParseCIDR(t, "104.131.0.0/33", "invalid CIDR address: 104.131.0.0/33")
}

func assertParseCIDR(t *testing.T, cidr string, expectFirst, expectLast uint32) {
	t.Helper()

	first, last, err := ParseCIDR(cidr)

	require.NoError(t, err)
	require.Equal(t, expectFirst, first)
	require.Equal(t, expectLast, last)
}

func assertExpectErrParseCIDR(t *testing.T, cidr string, msg string) {
	t.Helper()

	first, last, err := ParseCIDR(cidr)

	require.EqualError(t, err, msg)
	require.Equal(t, uint32(0), first)
	require.Equal(t, uint32(0), last)
}

func BenchmarkParseCIDR(b *testing.B) {
	for i := 0; i < b.N/3; i++ {
		_, _, _ = ParseCIDR("104.219.251.162/32")
		_, _, _ = ParseCIDR("104.131.0.0/16")
		_, _, _ = ParseCIDR("0.0.0.0/0")
	}
}
