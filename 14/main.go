package main

/* Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.*/

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan interface{})
	var numi int
	var numf float64
	var str string
	var b bool

	whatType(ch) //Использование рефлексии
	whatType(numi)
	whatType(numf)
	whatType(str)
	whatType(b)

	fmt.Println("Через Switch")
	whatTypeSwitch(ch)   //Использование switch data.(type)
	whatTypeSwitch(numi) // которая считывает любой тип данных
	whatTypeSwitch(numf) // и в зависимости от него, печатает этот самый тип
	whatTypeSwitch(str)
	whatTypeSwitch(b)

}

func whatTypeSwitch(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Printf("Тип данных int\n")
	case float64:
		fmt.Printf("Тип данных float\n")
	case string:
		fmt.Printf("Тип данных string\n")
	case bool:
		fmt.Printf("Тип данных bool\n")
	case chan interface{}:
		fmt.Printf("Тип данных chan\n")
	default:
		fmt.Printf("Неизвестный тип данных")
	}
}

func whatType(data interface{}) {

	fmt.Printf("Переменная - %v, Тип - %v\n", data, reflect.TypeOf(data))
}
