package bridge

import "fmt"

type mac struct {
	printer IPrinter
}

func (m *mac) SetPrinter(printer IPrinter) {
	m.printer = printer
}
func (m *mac) Print() {
	fmt.Println("mac: ready to print")
	m.printer.PrintFile()
}
