package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.*/

func main() {

	var WorkersNumber int

	fmt.Scan(&WorkersNumber) //Задаем кол-во воркеров

	jobs := make(chan int)         // Канал для получения работы
	results := make(chan struct{}) // канал для прекращения работы

	for i := 1; i <= WorkersNumber; i++ {
		go worker(i, jobs, results)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM) // Остановка по Ctrl + c

	go func() {
		<-sig // Получили сигнал, прекращаем чтение
		fmt.Println("\n Ctrl + c")
		close(results) //Закрываем каналы
		close(jobs)
	}()

	jobID := 1
	for {
		select {
		case <-results:
			fmt.Println("Основная горутина приостановила свою работу")
			return
		default:
			jobs <- jobID //Бесконечная отправка сообщений
			fmt.Printf("Отправлено действие %d\n", jobID)
			jobID++                     //каждый цикл ID увеличивается
			time.Sleep(1 * time.Second) //Для генерации помедленнее
		}

	}
}

func worker(id int, jobs <-chan int, results <-chan struct{}) {
	for {
		select {
		case job := <-jobs: // передаем ID работы
			fmt.Printf("Воркер %d начал работу %d\n", id, job)
		case <-results:
			fmt.Printf("Воркер %d закончил работу \n", id)
			return
		}
	}

}
