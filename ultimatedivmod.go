package piscine

func UltimateDivMod(div *int, mod *int) {
	x := *div
	*div = *div / *mod
	*mod = x % *mod
}
