package utils

import (
	"fmt"
	"strconv"
)

func HexWithOxPrefixToInt64(input string) int64 {
	output, _ := strconv.ParseInt(input, 0, 64)
	return output
}

func Int64ToHexWith0xPrefix(input int64) string {
	output := fmt.Sprintf("%#x", input)
	return output
}
