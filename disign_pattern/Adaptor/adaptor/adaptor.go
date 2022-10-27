package adaptor

import "fmt"

type usbAdaptor struct {
	winMachine *win
}

func NewUSBAdaptor(w *win) *usbAdaptor {
	fmt.Println("prepare an adaptor switch usb data to lighting")
	return &usbAdaptor{winMachine: w}
}

//适配器实现客户端接口，内部封装一个需要适配的对象引用
func (u *usbAdaptor) InsertIntoLightingPort() {
	u.winMachine.InsertIntoUSBPort()
}
