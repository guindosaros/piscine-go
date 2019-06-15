package piscine

import "math"

func IterativeFactorial(x int) int {
	if 0 > x {
		return 0
	} else if 0 == x || 1 == x {
		return 1
	} else {
		res := 1
		for i := 2; i <= x; i++ {
			res *= i
			if math.MaxInt32 < res {
				res = 0
				break
			}
		}
		return res
	}
}
