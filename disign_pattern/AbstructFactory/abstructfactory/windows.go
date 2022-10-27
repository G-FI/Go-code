package abstructfactory

import "fmt"

//product1_A 一组相关的对象
type WinButton struct{}

func (wb *WinButton) Press() {
	fmt.Println("WinButton pressed")
}

//product2_A
type WinCheckBox struct{}

func (wb *WinCheckBox) Check() {
	fmt.Println("WinCheckBox checked")
}

type WinFactory struct {
}

func (wf *WinFactory) CreateButton() Button {
	return &WinButton{}
}
func (wf *WinFactory) CreateCheckBox() CheckBox {
	return &WinCheckBox{}
}

func NewWinFactory() *WinFactory {
	return &WinFactory{}
}
