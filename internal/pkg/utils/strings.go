package utils

import (
	"strconv"
	"strings"
)

// Contains , check whether the number s contains the number sub
func Contains(s, sub uint64) bool {
	sStr := strconv.FormatUint(s, 10)
	subStr := strconv.FormatUint(sub, 10)
	return strings.Contains(sStr, subStr)
}
