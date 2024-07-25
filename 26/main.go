package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//
	str := "aabbcd"
	str1 := "avcd"
	str2 := "abCdefAaf"
	fmt.Println(compare(str))
	fmt.Println(compare(str1))
	fmt.Println(compare(str2))

}

func compare(s string) bool {

	uniqueChars := ""
	for _, char := range s { // Проходимся в цикле по всей входящей строке
		lowerChar := unicode.ToLower(char)                // Преобразования в нижний регистр
		if strings.ContainsRune(uniqueChars, lowerChar) { // проверка содержания текущего символа в строке
			return false
		}
		uniqueChars += string(lowerChar) // если символа еще нету, добавляем его
	}
	return true
}
