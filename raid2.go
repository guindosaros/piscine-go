package main

import (
	"reflect"
	"os"
	"fmt"
	"strings"
	piscine".."
)

func IsValideInput(args []string, grille, vertical, block *[][]string) bool {
	lon := 0
	if len(args) == 0|| len(args) == 1 || (piscine.Sqrt(len(args)) == 0 && len(args) > 1) {
		return false
	}
	for _,res:= range args {
		if reflect.TypeOf(res) != reflect.TypeOf("a") {
			return false
		}else{
			if lon == 0 {
				lon =len(res)
			}else{
				if lon != len(res) {
					return false
				}
			}
		}
	}
	for i,item:= range res {
		if string(item) != "." && strings.Count(res,string(item)) > 1 {
			return false
		}
		*(grille)[i] = pisicne.Split(res, "") 
	}
	for i:=0; i < lon; i++ {
		for j:=0; j < lon; j++ {
			*(vertical)[j][i] = *(grille)[i][j]
			*(block)[3(i/3)+(j/3)] = append(*(block)[3(i/3)+(j/3)],*(grille)[i][j])
		}
	}
	
	return IsValidGrille(grille, vertical, block, lon)
}

func printBoard(board [][]string) {
	fmt.Println("+-------+-------+-------+")
	for row,res:=range board {
		fmt.Print("| ")
		for col,item := range res{
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%s ", string(item))
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func hasDuplicate(grille *[][]string, n int) bool {
	for i:=0; i< n; i++ {
		res := ""
		for j:=0; j< n; i++ {
			res += string((*grille)[i][j])
		}
		for _,item:= range res {
			if string(item) != "." && strings.Count(res,string(item)) > 1 {
				return true
			}
		}
	}
	return false
}

func EmptyCell(grille *[][]string) bool {
	count := 0
	for i:=0;i < 9; i++ {
		for j:= 0; j<9; j++ {
			if grille[i][j] == "." {
				count++
			}
		}
	}
	return count == 0
}

func IsValidGrille(grille, vertical, block *[][]string, lon int) bool {
	if hasDuplicate(&grille, lon) || hasDuplicate(&block, lon) || hasDuplicate(&vertical, lon) {
		return false
	}
	return true
}

func Backtracking(grille, vertical, block *[][]string) bool {
	if !EmptyCell(&grille){
		return true
	}
	for i:=0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*grille)[i][j] == '.' {
				for candidate := 9; candidate >= 1; candidate-- {
					(*grille)[i][j] = string(candidate)
					if IsValidGrille(grille, vertical,block,9) {
						if Backtracking(grille, vertical, block) {
							return true
						}
						(*grille)[i][j] =  string(0)
					} else {
						(*grille)[i][j] = string(0)
					}
				}
				return false
			}
		}
	}
	return false
}

func main(){
	args:= os.Args[1:]
	row, col, block [][]string 
	if len(args) == 0 {
		fmt.Println("Error")
	}else{
		if IsValideInput(args, &row, &col, &block){
			if !Backtracking(&row,&col, &block) {
				fmt.Println("Error")
			}else{
				printBoard(row)
			}
		}else{
			fmt.Println("Error")
		}
	}
}
