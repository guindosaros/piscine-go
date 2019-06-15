package piscine

func IterativePower(nb, power int) int {
	if power < 0 {
		return 0
	} else if 1 == power {
		return nb
	} else if 0 == power {
		return 1
	} else {
		res := 1
		for i := 0; i < power; i++ {
			res *= nb
		}
		return res
	}
}
