package main

/* К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации. */

/*
При создании большой строки, мы выделяем для неё память,
после, присваивая глобальной переменной ссылку на эту же большую строку
Планировщик, не чистит память используемую глобальной переменной, т.к
juststring ссылается на v, но и его тоже чистить не будет, что приведет к утечке памяти
*/
// var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString := make([]rune, 100)
	copy(justString, v[:100]) // копируем а не присваиваем
}
func main() {
	someFunс()

}
