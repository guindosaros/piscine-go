package piscine

func Compact(ptr *[]string, length int) int {
	count:=0
	for i:=0;i<length;i++{
		if (*ptr)[i]!=" "{
			count++
		}
	}
	return count
}