package piscine

func Fibonacci(x int) int {
	if x < 0 {
		return -1
	} else if 0 == x {
		return 0
	} else if 1 == x || 2 == x {
		return 1
	} else {
		return Fibonacci(x-1) + Fibonacci(x-2)
	}
}
