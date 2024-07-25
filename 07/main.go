package main

import (
	"fmt"
	"sync"
)

type storage struct {
	m  map[int]int
	mu sync.Mutex // замок
}

func (s *storage) write(i int) {
	s.mu.Lock() // закрытие
	s.m[0] = i
	defer s.mu.Unlock() // открытие

}

func main() {
	var wg sync.WaitGroup

	wg.Add(1) //т.к знаем точное кол-во горутин, добавляю 1 к счетчику
	s := storage{
		m: make(map[int]int),
	}

	go func() {
		defer wg.Done() // горутина завершается
		for i := 0; i < 1000000; i++ {
			s.write(i)

		}

	}()

	wg.Wait()
	fmt.Printf("%d", s)
}
