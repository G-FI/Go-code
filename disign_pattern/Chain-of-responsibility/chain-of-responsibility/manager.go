package chain_of_responsibility

import "fmt"

type manager struct {
	nextHandler iHandler
}

func NewManager() *manager {
	return &manager{nextHandler: nil}
}

func (m *manager) SetNext(handler iHandler) {
	m.nextHandler = handler
}

func (m *manager) Handle() {
	fmt.Println("manager is handling")
	if m.nextHandler != nil {
		m.nextHandler.Handle()
	}
}
