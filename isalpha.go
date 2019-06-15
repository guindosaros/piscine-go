package piscine

import "regexp"

func IsAlpha(str string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	return re.MatchString(str)
}
