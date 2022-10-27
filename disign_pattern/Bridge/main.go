package main

import "bridge/bridge"

func main() {
	hp, _ := bridge.NewPrinter("hp")
	epson, _ := bridge.NewPrinter("epson")

	win, _ := bridge.NewComputer("win", hp)
	win.Print()
	win.SetPrinter(epson)
	win.Print()

	mac, _ := bridge.NewComputer("mac", hp)
	mac.Print()
	mac.SetPrinter(epson)
	mac.Print()

	//timerPrinter就是独立扩展的一个具体实现部分，通过在抽象部分中组合新的具体对象来，利用对象的方法来扩展功能
	timerPrinter, _ := bridge.NewPrinter("timerPrinter")
	mac.SetPrinter(timerPrinter)
	mac.Print()
}
