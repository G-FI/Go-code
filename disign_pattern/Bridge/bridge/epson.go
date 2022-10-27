package bridge

import "fmt"

type epson struct {
}

func (e *epson) PrintFile() {
	fmt.Println("epson is printing file")
}
