package chain_of_responsibility

import "fmt"

type checker struct {
	nextHandler iHandler
}

func (c *checker) SetNext(handler iHandler) {
	c.nextHandler = handler
}

func (c *checker) Handle() {
	fmt.Println("checker is handler")
	if c.nextHandler != nil {
		c.nextHandler.Handle()
	}
}

func NewChecker() *checker {
	return &checker{nextHandler: nil}
}
