package main

import (
	"context"
	"fmt"
	"time"
)

/*Реализовать все возможные способы остановки выполнения горутины.*/

/* Способы остановки выполнения горутины
1. Завершить Main быстрее чем завершатся горутины
2.  Создать условие при котором, если в канал поступит сигнал, произойдет return
3. Контекст
*/

func main() {
	quit := make(chan bool)
	go workerbool(quit)
	close(quit)
	time.Sleep(time.Second * 1)
	ctx, cancel := context.WithCancel(context.Background())
	go workercontext(ctx)
	cancel()
	time.Sleep(time.Second * 1)

}

func workercontext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершен через контекст")
			return
		default:
			fmt.Println("Работает контекст")
		}
	}

}
func workerbool(stop chan bool) {
	for {

		select {
		case <-stop: //Если в канал подается значение, то срабатывает этот кейс
			fmt.Println("Выход через канал")
			return
		default:
			fmt.Println("Сообщение")
		}

	}
}
