package chain_of_responsibility

import "fmt"

type cashier struct {
	nextHandler iHandler
}

func (c *cashier) SetNext(handler iHandler) {
	c.nextHandler = handler
}

func (c *cashier) Handle() {
	fmt.Println("cashier is handling")
	if c.nextHandler != nil {
		c.nextHandler.Handle()
	}
}

func NewCashier() *cashier {
	return &cashier{nextHandler: nil}
}
