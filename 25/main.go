package main

import (
	"fmt"
	"time"
)

func main() {

	sleep(5 * time.Second)

	fmt.Println("Wassup")
}

func sleep(x time.Duration) {
	<-time.After(x) // After ожидает x и после возвращает канал с текущим временем

}
