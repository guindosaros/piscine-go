package piscine

import "regexp"

func IsNumeric(str string) bool {
	re := regexp.MustCompile("^[0-9_]*$")
	return re.MatchString(str)
}
