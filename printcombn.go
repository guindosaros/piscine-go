package piscine

import "github.com/01-edu/z01"

func print(i int) {
	z01.PrintRune(rune(i) + '0')
}

func show(n int, tab [9]int, max [9]int) {
	i := 0
	for i < n {
		print(tab[i])
		i++
	}
	if tab[0] != max[0] {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func combinaison1() {
	tab := [9]int{0}
	max := [9]int{9}
	for tab[0] <= max[0] {
		show(1, tab, max)
		tab[0]++
	}
}

func PrintCombN(n int) {
	tab := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	max := [9]int{}

	if n == 1 {
		combinaison1()
	} else {
		i := n - 1
		j := 9
		for i >= 0 {
			i--
			j--
			max[i] = j
		}
		i = n - 1
		for tab[0] < max[0] {
			if tab[i] != max[i] {
				show(n, tab, max)
				tab[i]++
			}
			if tab[i] == max[i] {
				if tab[i-1] != max[i-1] {
					show(n, tab, max)
					tab[i-1]++
					j = i
					for j < n {
						tab[j] = tab[j-1] + 1
						j++
					}
					i = n - 1
				}
			}
			for tab[i] == max[i] && tab[i-1] == max[i-1] && i > 1 {
				i--
			}
		}
		show(n, tab, max)
	}
	z01.PrintRune('\n')
}
