package bridge

import "fmt"

//抽象部分
type IComputer interface {
	Print()
	SetPrinter(IPrinter)
}

func NewComputer(name string, printer IPrinter) (IComputer, error) {
	if name == "mac" {
		return &mac{printer: printer}, nil
	} else if name == "win" {
		return &win{printer: printer}, nil
	}
	return nil, fmt.Errorf("not support plantform")
}
