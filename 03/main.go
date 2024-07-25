package main

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

import (
	"fmt"
	"sync"
)

func main() {

	arr := []int{2, 4, 6, 8, 10}
	var sum int
	var wg sync.WaitGroup
	for i := 0; i < len(arr); i++ {
		wg.Add(1)        // 1 горутина, 1 инкрементация счетчика
		go func(i int) { //анонимная функция в горутине, высчитывающая сумму квадратов массива
			defer wg.Done() // инкрементация
			square := arr[i] * arr[i]
			sum += square
		}(i)
	}
	wg.Wait() // Ожидаем пока выполнятся все горутины
	fmt.Println("Sum of squares:", sum)
}
