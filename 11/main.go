package main

import "fmt"

/* Реализовать пересечение двух неупорядоченных множеств. */
func main() {

	var group = []string{"cat", "cat", "dog", "cat", "tree"}
	var anotherGroup = []string{"cat", "zebra", "bird"}
	set := make(map[string]struct{}, len(group)) // ключи string с множеством struct
	anotherSet := make(map[string]struct{}, len(anotherGroup))
	for _, v := range anotherGroup {
		anotherSet[v] = struct{}{}
	}
	for _, v := range group { // добавляем строки в множество
		set[v] = struct{}{} // использование пустой структуры означает что нужно только наличие ключа
		//если ключ существует, ничего не меняется
	}
	//итерация по ключам карты
	for key := range set {
		if _, ok := anotherSet[key]; ok {
			fmt.Print(key, " ")
		}
	}
}
