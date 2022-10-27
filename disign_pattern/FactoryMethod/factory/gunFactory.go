package factory

import "fmt"

func CreateGun(name string) (IGun, error) {
	if name == "ak47" {
		return NewAk47(), nil
	} else if name == "m416" {
		return NewM416(), nil
	}
	return nil, fmt.Errorf("wrong gun type")
}
