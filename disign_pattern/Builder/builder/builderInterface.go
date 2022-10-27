package builder

import "fmt"

type Builder interface {
	Reset()
	SetSeats()
	SetEngin()
}

func NewBuilder(name string) (Builder, error) {
	if name == "car" {
		return newCarBuilder(), nil
	} else if name == "carmanual" {
		return newCarManualBuilder(), nil
	}
	return nil, fmt.Errorf("not support builder")
}
