package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
//массива, во второй — результат операции x*2, после чего данные из второго
//канала должны выводиться в stdout.

func writenums(ch1 chan<- int, x []int) {
	for _, v := range x {
		ch1 <- v // записываем в ch1 числа
	}
	close(ch1) // закрытие канала
}

func doublenums(ch2 chan<- int, ch1 <-chan int) { // c2 - записывает, ch1 - читает
	for num := range ch1 {

		ch2 <- num * 2 // записываем значения в канал
	}
	close(ch2) // закрытие канала
}
func main() {
	var x = []int{2, 5, 8, 10}
	ch1 := make(chan int) // инициализация каналов
	ch2 := make(chan int)

	go writenums(ch1, x)
	go doublenums(ch2, ch1)
	for result := range ch2 {
		fmt.Println(result) // вывод x *2
	}

}
