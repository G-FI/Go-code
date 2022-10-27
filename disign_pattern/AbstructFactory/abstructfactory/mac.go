package abstructfactory

import "fmt"

//product1_B
type MacButton struct{}

func (m *MacButton) Press() {
	fmt.Println("MacButton pressed")
}

//prouct2_B
type MacCheckBox struct{}

func (m *MacCheckBox) Check() {
	fmt.Println("MacCheckBox checked")
}

type MacFactory struct{}

func (mf *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (mf *MacFactory) CreateCheckBox() CheckBox {
	return &MacCheckBox{}
}

func NewMacFactory() *MacFactory {
	return &MacFactory{}
}
