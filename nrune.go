package piscine

func NRune(s string, x int) rune {
	res := []rune(s)
	return res[x-1]
}
