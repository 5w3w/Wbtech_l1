package main

/* Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode. */

import (
	"fmt"
	"strings"
)

func main() {

	var str string = "Привет"
	reverseString(str)
}

func reverseString(s string) {

	newS := strings.Split(s, "")

	for i := len(newS) - 1; i >= 0; i-- { //Считаем с конца строки, пока не декрементируемся до нуля.
		fmt.Printf(newS[i])
	}
}
