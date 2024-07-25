package main

import (
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func main() {
	var num int64
	var i uint
	fmt.Printf("Число - ")
	fmt.Scanf("%d", &num)
	fmt.Printf("Какой бит меняем? - ")
	fmt.Scanf("%d", &i)

	fmt.Printf("Битовое представление %d - %b\n", num, num)

	num = setbit(num, i)
	fmt.Printf("Устанавливаем %d бит в 1 или 0: %b = %d \n", i, num, num)

}

func setbit(num int64, i uint) int64 {
	return (num ^ (1 << i)) // Создание маски и использование XOR
}
