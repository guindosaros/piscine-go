package piscine

func BasicJoin(strs []string) string {
	concat := ""
	for _, res := range strs {
		concat += res
	}
	return concat
}
