package chain_of_responsibility

import "fmt"

type clerk struct {
	nextHandler iHandler
}

func (c *clerk) SetNext(handler iHandler) {
	c.nextHandler = handler
}

func (c *clerk) Handle() {
	fmt.Println("clerk is handling")
	if c.nextHandler != nil {
		c.nextHandler.Handle()
	}
}

func NewClerk() *clerk {
	return &clerk{nextHandler: nil}
}
