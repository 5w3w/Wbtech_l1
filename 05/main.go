// /*Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться. */

// /*
// 2 канала, 1 на чтение, 2ой на запись
// формируем данные
// отправляем эти данные N секунд
// отправляем их в 1 канал, из другого читаем
// закрываем каналы после N секунд
// Работы программы завершена
// */
// package main
package main

import (
	"fmt"
	"time"
)

func main() {
	var duration int
	fmt.Scan(&duration)
	chan1 := make(chan int)
	chan2 := make(chan int)
	done := make(chan bool)

	go generate_num(chan1, duration)
	go count(chan1, chan2)

	go func() {
		for num := range chan2 {
			fmt.Println(num) // Читаем и печатаем значения из второго канала
		}
	}()

	time.Sleep(time.Duration(duration) * time.Second) // Ждем N секунд
	close(done)                                       // Закрываем канал done, чтобы завершить программу
}

func count(ch1 chan int, ch2 chan int) {
	for num := range ch1 {
		ch2 <- num
	}
	close(ch2) // Закрываем второй канал после чтения всех данных
}

func generate_num(ch chan<- int, timer int) {
	for i := 0; i < timer; i++ {
		ch <- i
		time.Sleep(1 * time.Second) // Пауза в 1 секунду
	}
	close(ch) // Закрываем канал после отправки всех данных
}
