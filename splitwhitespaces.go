package piscine

import "strings"

func SplitWhiteSpaces(str string) []string {
	tab := strings.Fields(str)
	return tab
}
