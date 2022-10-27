package abstructfactory

type AbsFactory interface {
	CreateButton() Button
	CreateCheckBox() CheckBox
}
