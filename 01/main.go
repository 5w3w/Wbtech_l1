package main

import "fmt"

// Структура(Родительская) Human
type Human struct {
	Name string
	age  int
}

// Структура(Дочерняя) Action со встроенной структурой Human
type Action struct {
	Human
}

// Метод Say() структуры Human
func (h *Human) Say() {
	fmt.Printf("Hello, my name is %v\n", h.Name)
}

func main() {

	Hum1 := &Human{ //Определение переменной типа Human, к которой применяется метод Say()
		"Nicolas",
		32,
	}

	Mv1 := Action{ // Определение переменной типа Action, которая наследует свойства структуры Human
		Human{
			"Roma",
			12,
		},
	}
	Hum1.Say() // Вызов метода Say() для переменной Hum1, т.к она определенна типом Human к которой относится метод Say

	Mv1.Say() //Вызов метода Say() для переменной Mv1, т.к она наследует свойства структуры Human которая,
	// в свою очередь, использует метод Say()
}
