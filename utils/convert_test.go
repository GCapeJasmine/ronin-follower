package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHexWithOxPrefixToInt64(t *testing.T) {
	input := "0x4d2"
	output := HexWithOxPrefixToInt64(input)
	require.Equal(t, int64(1234), output)
}

func TestInt64ToHexWith0xPrefix(t *testing.T) {
	input := int64(1234)
	output := Int64ToHexWith0xPrefix(input)
	require.Equal(t, "0x4d2", output)
}
