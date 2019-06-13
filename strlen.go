package piscine

import "unicode/utf8"

func StrLen(str string) int {
              a := utf8.RuneCountInString(str)
              return a
}