package main

/* Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.*/
import "fmt"

func main() {
	var group = []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{}, len(group)) // ключи string с множеством struct

	for _, v := range group { // добавляем строки в множество
		set[v] = struct{}{} // использование пустой структуры означает что нужно только наличие ключа
		//если ключ существует, ничего не меняется
	}
	//итерация по ключам карты
	for key := range set { // выводим множество
		fmt.Println(key)
	}
}
