package incidr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIn(t *testing.T) {
	require.Equal(t, false, In(nil, 1))
	require.Equal(t, true, In([]Range{
		{
			First: 0,
			Last:  255,
		},
	}, 1))
}
