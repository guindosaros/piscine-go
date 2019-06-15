package piscine

func LastRune(s string) rune {
	res := []rune(s)
	return res[len(res)-1]
}
