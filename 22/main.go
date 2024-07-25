package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	a := big.NewFloat(math.Pow(2, 20) * 2) // Библиотека math/big
	b := big.NewFloat(math.Pow(2, 20) * 3) // для работы арифметикой произвольной точности

	fmt.Println(a.Add(a, b))
	fmt.Println(a.Sub(a, b))
	fmt.Println(a.Mul(a, b))
	fmt.Println(a.Quo(a, b))
	/*
		Add, Sub, Mul, Quo методы big
	*/

}
