package incidr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseIPv6CIDR(t *testing.T) {
	assertParseIPv6CIDR(t, "0:0:0:1::0/96", IPv6{1, 0}, IPv6{1, 0xffffffff})
	assertParseIPv6CIDR(t, "0:0:0:1::0/96", IPv6{1, 0}, IPv6{1, 0xffffffff})
	assertParseIPv6CIDR(t, "0:0:0:1::0/64", IPv6{1, 0}, IPv6{1, 0xffffffffffffffff})

	assertParseIPv6CIDR(t, "1:2:3:4::0/32", IPv6{0x0001000200000000, 0}, IPv6{0x00010002ffffffff, 0xffffffffffffffff})
}

func assertParseIPv6CIDR(t *testing.T, cidr string, expectFirst, expectLast IPv6) {
	t.Helper()

	var first, last, err = ParseIPv6CIDR(cidr)

	require.NoError(t, err)
	require.Equal(t, expectFirst, first)
	require.Equal(t, expectLast, last)
}

func BenchmarkParseIPv6CIDR(b *testing.B) {
	for i := 0; i < b.N/5; i++ {
		_, _, _ = ParseIPv6CIDR("1:2:3:4::0/128")
		_, _, _ = ParseIPv6CIDR("1:2:3:4::0/96")
		_, _, _ = ParseIPv6CIDR("1:2:3:4::0/64")
		_, _, _ = ParseIPv6CIDR("1:2:3:4::0/32")
		_, _, _ = ParseIPv6CIDR("1:2:3:4::0/0")
	}
}
