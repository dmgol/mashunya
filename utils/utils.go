package utils

import "strconv"

func ParseUintRef(str string, out *uint) bool {
	if value, err := strconv.ParseUint(str, 10, 32); err == nil {
		*out = uint(value)
		return true
	}
	return false
}

func ParseUint(p string) (value uint, success bool) {
	success = ParseUintRef(p, &value)
	return
}
