package piscine

import "math"

func RecursiveFactorial(x int) int {
	if x < 0 || x > math.MaxInt32 {
		return 0
	} else if 0 == x || 1 == x {
		return 1
	} else {
		res := x*RecursiveFactorial(x-1) > math.MaxInt32
		if res {
			return 0
		} else {
			return x * RecursiveFactorial(x-1)
		}
	}
}
