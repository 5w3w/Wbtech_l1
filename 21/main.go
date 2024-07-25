package main

import "fmt"

type Target interface {
	Print(string) string
}

type OldPrinter struct{}
type NewPrinter struct{}

type PrintAdapter struct {
	newPrinter *NewPrinter
}

func (oldP *OldPrinter) Print(msg string) string {
	return "Old printer " + msg
}

func (newP *NewPrinter) newPrint(msg string) string { // newPrint не соответсвует интерфейсу Target
	return "New printer " + msg
}
func (printA *PrintAdapter) Print(msg string) string { //Адаптирует метод Print для newPrinter
	return printA.newPrinter.newPrint(msg)
}

func main() {
	var printer Target // Используем переменную printer типа Target

	printer = &OldPrinter{} //Ссылаемся на OldPrinter который реализует метод Print
	fmt.Println(printer.Print("hee-heee"))

	newPrinter := &NewPrinter{}
	printer = &PrintAdapter{newPrinter: newPrinter}
	/*
		Ссылаемся на PrintAdapter, который содержит ссылку на NewPrint и
		адаптирует метод Print интерфейса Target к методу PrintNew нового принтера
	*/
	fmt.Println(printer.Print("hee-heeeee new printer "))
	/*
		Когда клиент вызывает Print на PrinterAdapter,
		метод Print адаптера вызывает метод PrintNew у NewPrinter.
	*/
}
