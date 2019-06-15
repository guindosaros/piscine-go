package piscine

import "unicode"

func IsUpper(str string) bool {
	for _, res := range str {
		if !unicode.IsUpper(res) {
			return false
		}
	}
	return true
}
