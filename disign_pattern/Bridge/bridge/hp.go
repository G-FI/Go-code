package bridge

import "fmt"

type hp struct {
}

func (h *hp) PrintFile() {
	fmt.Println("hp is printing file")
}
