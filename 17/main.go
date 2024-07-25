package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	var arr = []int{1, 3, 4, 8, 12, 20, 34}
	fmt.Println(binarysearch(arr, 34))
}

func binarysearch(nums []int, target int) interface{} {

	left, right := 0, len(nums) // определяем меньший и больший элементый массива

	for left < right {
		midindex := int((left + right) / 2) //Дробление массива на пополам
		midvalue := nums[midindex]          // передаем значение числа по середине
		if midvalue == target {             // если число совпадает, возвращаем индекс найденного числа
			return left
		} else if midvalue < target { // определяется новый массив в зависимости от равенства
			left = midindex + 1 // ищем с правой части массива

		} else {
			right = midindex - 1 //ищем с левой
		}
	}
	return fmt.Sprint("Число не входит в массив") //Возвращаем текст если числа нет в массиве
}

/*  {1, 3, 4, 8, 12, 20, 34}, target(t) = 34
midindex := (0 + 6) / 2 = 3
midvalue = nums[3] = 8
8 < 34
left = midindex +1 (3 +1)
Формируется новая область поиска
{12, 20, 34} - индексы не меняются (4, 5, 6)
ищем новый средний элемент
midindex := (4 + 6) / 2 = 5
midvalue := nums[5] = 20
20<34
left = midindex + 1 = 5 + 1 = 6
еще раз формируем новую область
midindex := (6 + 6) / 2 = 6
midvalue := nums[6] = 34
midvalue == target (34 == 34)
{
возвращаем индекс
}
*/
