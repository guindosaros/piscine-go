package piscine

import "sort"

func Max(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}