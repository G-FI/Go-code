package bridge

import "fmt"

type win struct {
	printer IPrinter
}

func (w *win) Print() {
	fmt.Println("win ready to print")
	w.printer.PrintFile()
}
func (w *win) SetPrinter(printer IPrinter) {
	w.printer = printer
}
