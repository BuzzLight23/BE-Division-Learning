package utils

import (
	"strconv"
)

func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}