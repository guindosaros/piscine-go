package piscine

import "strings"

func Capitalize(s string) string {
	res := strings.Title(ToLower(s))
	indx := Index(s, "_")
	if indx == -1 {
		return res
	} else {
		res = strings.Replace(res, string(res[indx+1]), ToUpper(string(res[indx+1])), indx+1)
		return res
	}
}
