package bridge

import "fmt"

type IPrinter interface {
	PrintFile()
	//扩展
}

func NewPrinter(name string) (IPrinter, error) {
	if name == "epson" {
		return &epson{}, nil
	} else if name == "hp" {
		return &hp{}, nil
	} else if name == "timerPrinter" {
		return newTimerPrinter("2022-3-24"), nil
	}
	return nil, fmt.Errorf("not support printer type")
}
