package main

//Удалить i-ый элемент из слайса.
import (
	"fmt"
	"slices"
)

func main() {

	arr := []int{1, 3, 4, 5, 6}
	arrCopy := make([]int, len(arr)) // Изменения элементов массива внутри функции, отражаются на исходном срезе
	copy(arrCopy, arr)
	fmt.Println("first func ", deleteElem(arr, 1))       //Изменения оригинального среза
	fmt.Println("Second func ", deleteElem2(arrCopy, 1)) //Т.к есть полная копия исходного массива, то вывод не ломается
}

/*
В go когда мы делаем срез [:i] мы не берем сам i в массиве,
поэтому, для удаления нужно начать с i+1
*/
func deleteElem2(slice []int, i int) []int {
	j := i
	slice = append(slice[:i], slice[j+1:]...) //
	// clear(slice)
	return slice
}

// Используется пакет slice (go 1.21), функция Delete для удаления элементов массива с i по j

func deleteElem(slice []int, i int) []int {
	i = 1
	j := i
	newSlice := slices.Delete(slice, i, j+1)
	return newSlice
}
