package main

/* Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow». */

import (
	"fmt"
	"strings"
)

func main() {
	type1 := "cat dog snow"
	Reverse(type1)
	// fmt.Println(len(type1))
}

/*
Подаем строку
Разделяем её
Перемещаем
*/
func Reverse(s string) {
	news := strings.Split(s, " ")

	for i := len(news) - 1; i >= 0; i-- { ////Считаем с конца строки, пока не декрементируемся до нуля.
		fmt.Print(news[i], " ")
	}

}
