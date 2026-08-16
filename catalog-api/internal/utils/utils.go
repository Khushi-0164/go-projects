package utils

import "strconv"

func ParseUint(s string) uint {
	n, _ := strconv.ParseUint(s, 10, 64)
	return uint(n)
}
