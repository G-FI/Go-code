package adaptor

import "fmt"

type win struct {
}

func NewWin() *win {
	return &win{}
}

func (w *win) InsertIntoUSBPort() {
	fmt.Println("win: insert into a usb port")
}
