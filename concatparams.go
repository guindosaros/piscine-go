package piscine

func ConcatParams(args []string) string {
	str := ""

	for i,res:= range args {
		str += string(res)
		if i != len(args) - 1 {
			str +=  "\n"
		}
	}
	return str
}
