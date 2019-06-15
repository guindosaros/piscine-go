package piscine

import "math"

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	} else if 2 == nb || 3 == nb {
		return true
	} else {
		limit := math.Sqrt(float64(nb))
		for i := 2; i <= int(limit); i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
}
