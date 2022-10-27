package chain_of_responsibility

import (
	"fmt"
)

type receptor struct {
	nextHandler iHandler
}

func NewReceptor() *receptor {
	return &receptor{nextHandler: nil}
}

func (r *receptor) SetNext(handler iHandler) {
	r.nextHandler = handler
}

func (r *receptor) Handle() {
	if r.nextHandler == nil {
		fmt.Println("receptor is handling")
		return
	}
	switch r.nextHandler.(type) {
	case *manager:
		fmt.Println("receptor is handling VIP customer")
	default:
		fmt.Println("receptor is handling ordinary customer")
	}
	r.nextHandler.Handle()
}
