package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка
*/
func main() {
	var arr = []int{1, 8, 3, 222, 13, 12, 20, 34}
	fmt.Println(quicksort(arr, 0, len(arr)-1))
}

func quicksort(arr []int, low, high int) []int {

	if low < high {
		var p int
		arr, p = patrition(arr, low, high)
		arr = quicksort(arr, low, p-1)
		arr = quicksort(arr, p+1, high)
	}
	return arr

}

func patrition(arr []int, low int, high int) ([]int, int) {
	pivot := arr[high]            // 34 - опорный элемент с индексом high
	i := low                      // 0
	for j := low; j < high; j++ { // Специально определяем j := low, так мы, при желании сможет отсортировать конкретный участок массива
		if arr[j] < pivot { // Если элемент массива меньше опорного (34)
			arr[i], arr[j] = arr[j], arr[i] // то перемещаем его
			i++                             //1 < 34, тогда остается на своем месте
			//если не соблюдается условие то и отсутствет инкрементация
		}

	}
	arr[i], arr[high] = arr[high], arr[i] // переставляется опорный элемент
	// и после перестановки, все, что левее него, меньше, а правее больше

	return arr, i
}
