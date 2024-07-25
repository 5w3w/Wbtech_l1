package main

import (
	"fmt"
	"unsafe"
)

type test struct{}

func main() {

	v1 := test{}
	v2 := struct{}{}
	fmt.Printf("Size of v1 - %v\n", unsafe.Sizeof(v1)) //0
	fmt.Printf("Size of v2 - %v\n", unsafe.Sizeof(v2)) //0

}
