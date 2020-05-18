package test

import (
	"github.com/YaroslavPodorvanov/go-incidr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCidrList_Contains(t *testing.T) {
	var cidrlist = incidr.NewCidrList()

	require.Equal(t, false, cidrlist.Contains(1))
	require.Equal(t, false, cidrlist.Contains(1<<24|1<<16|1<<8+1))
	require.Equal(t, false, cidrlist.Contains(2<<24|1<<16|1<<8+2))

	cidrlist.Update(incidr.RangeSequenceByCIDRs([]string{"1.1.1.0/24", "2.1.1.0/24"}))

	require.Equal(t, false, cidrlist.Contains(1))
	require.Equal(t, true, cidrlist.Contains(1<<24|1<<16|1<<8+1))
	require.Equal(t, true, cidrlist.Contains(2<<24|1<<16|1<<8+2))
}
