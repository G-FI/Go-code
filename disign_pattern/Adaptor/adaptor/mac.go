package adaptor

import "fmt"

type mac struct {
}

func NewMac() *mac {
	return &mac{}
}

func (m *mac) InsertIntoLightingPort() {
	fmt.Println("Mac: insert into a lighting Port")
}
