package main

import (
	"adaptor/adaptor"
	"fmt"
)

type client struct {
}

func (c *client) Insert(com adaptor.IComputer) {
	fmt.Println("clint insert a lighting connector into a computer")
	com.InsertIntoLightingPort()
}

func main() {
	mac := adaptor.NewMac()
	win := adaptor.NewWin()
	c := &client{}

	c.Insert(mac)
	//c.Insert(win) win没有实现IComputer接口

	fmt.Println("===============")
	a := adaptor.NewUSBAdaptor(win)
	c.Insert(a)
}
