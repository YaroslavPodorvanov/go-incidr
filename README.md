# InCidr [![Build Status](https://travis-ci.org/YaroslavPodorvanov/go-incidr.svg?branch=master)](https://travis-ci.org/YaroslavPodorvanov/go-incidr)
Lookup IP in CIDR list

### Examples
```golang
package test

import (
	"github.com/YaroslavPodorvanov/go-incidr"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestIpv4CidrList_Contains(t *testing.T) {
	var ipv4Cidrlist = incidr.NewIPv4CidrList()

	require.Equal(t, false, ipv4Cidrlist.Contains(1))
	require.Equal(t, false, ipv4Cidrlist.Contains(1<<24|1<<16|1<<8+1))
	require.Equal(t, false, ipv4Cidrlist.Contains(2<<24|1<<16|1<<8+2))

	ipv4Cidrlist.Update(incidr.RangeSequenceByIPv4CIDRs([]string{"1.1.1.0/24", "2.1.1.0/24"}))

	require.Equal(t, false, ipv4Cidrlist.Contains(1))
	require.Equal(t, true, ipv4Cidrlist.Contains(1<<24|1<<16|1<<8+1))
	require.Equal(t, true, ipv4Cidrlist.Contains(2<<24|1<<16|1<<8+2))
}

func TestCidrList_Contains(t *testing.T) {
	var (
		ipv4Cidrlist = incidr.NewIPv4CidrList()
		ipv6Cidrlist = incidr.NewIPv6CidrList()
		cidrlist     = incidr.NewCidrList(ipv4Cidrlist, ipv6Cidrlist)
	)

	require.Equal(t, false, cidrlist.Contains(net.ParseIP("0.0.0.1")))
	require.Equal(t, false, cidrlist.Contains(net.ParseIP("1:2:3:4::1")))
	require.Equal(t, false, cidrlist.Contains(net.ParseIP("1:2:3:5::1")))

	ipv4Cidrlist.Update(incidr.RangeSequenceByIPv4CIDRs([]string{"0.0.0.1/24"}))
	ipv6Cidrlist.Update(incidr.RangeSequenceByIPv6CIDRs([]string{"1:2:3:4::1/118"}))

	require.Equal(t, true, cidrlist.Contains(net.ParseIP("0.0.0.1")))
	require.Equal(t, true, cidrlist.Contains(net.ParseIP("1:2:3:4::1")))
	require.Equal(t, false, cidrlist.Contains(net.ParseIP("1:2:3:5::1")))
}
```

### Alternatives
* [github.com/yl2chen/cidranger](https://github.com/yl2chen/cidranger)
* [github.com/c-robinson/iplib](https://github.com/c-robinson/iplib)
