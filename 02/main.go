package main

/* Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
import (
	"fmt"
	"sync"
)

func main() {

	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, v := range arr { //итерируемся по значениям слайса
		wg.Add(1)         // инкрементируем счётчик waitgroup
		go square(v, &wg) //передаем декрементируемую waitgroup и значение квадрата числа
	}
	wg.Wait() // ожидание выполнения всех горутин
}

func square(num int, wg *sync.WaitGroup) {
	defer wg.Done() //Декрементируем счетчик
	fmt.Printf("%d * %d = %d \n", num, num, num*num)

}
