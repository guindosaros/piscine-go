package piscine

import "unicode"

func IsPrintable(str string) bool {
	for _, res := range str {
		if !unicode.IsPrint(res) {
			return false
		}
	}
	return true
}
