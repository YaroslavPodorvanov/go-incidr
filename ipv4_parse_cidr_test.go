package incidr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseIPv4CIDR(t *testing.T) {
	assertParseIPv4CIDR(t, "104.131.0.0/16", 0x68830000, 0x6883ffff)
	assertParseIPv4CIDR(t, "104.219.251.162/32", 0x68dbfba2, 0x68dbfba2)
	assertParseIPv4CIDR(t, "1.2.64.0/18", 0x1024000, 0x1027fff)
	assertParseIPv4CIDR(t, "0.0.0.0/0", 0x0, 0xffffffff)
}

func TestParseInvalidCIDRAddress(t *testing.T) {
	assertExpectErrParseIPv4CIDR(t, "104.131.0.0/", "invalid CIDR address: 104.131.0.0/")
	assertExpectErrParseIPv4CIDR(t, "104.131.0.0/33", "invalid CIDR address: 104.131.0.0/33")
}

func assertParseIPv4CIDR(t *testing.T, cidr string, expectFirst, expectLast uint32) {
	t.Helper()

	var first, last, err = ParseIPv4CIDR(cidr)

	require.NoError(t, err)
	require.Equal(t, expectFirst, first)
	require.Equal(t, expectLast, last)
}

func assertExpectErrParseIPv4CIDR(t *testing.T, cidr string, msg string) {
	t.Helper()

	first, last, err := ParseIPv4CIDR(cidr)

	require.EqualError(t, err, msg)
	require.Equal(t, uint32(0), first)
	require.Equal(t, uint32(0), last)
}

func BenchmarkParseIPv4CIDR(b *testing.B) {
	for i := 0; i < b.N/3; i++ {
		_, _, _ = ParseIPv4CIDR("104.219.251.162/32")
		_, _, _ = ParseIPv4CIDR("104.131.0.0/16")
		_, _, _ = ParseIPv4CIDR("0.0.0.0/0")
	}
}
