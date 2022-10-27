package main

import (
	chain_of_responsibility "Chain-of-responsibility/chain-of-responsibility"
	"fmt"
)

func main() {
	receptor := chain_of_responsibility.NewReceptor()
	clerk := chain_of_responsibility.NewClerk()
	checker := chain_of_responsibility.NewChecker()
	manager := chain_of_responsibility.NewManager()
	cashier := chain_of_responsibility.NewCashier()

	fmt.Println("ordinary customer chain:")
	receptor.SetNext(clerk)
	clerk.SetNext(checker)
	checker.SetNext(manager)
	manager.SetNext(cashier)
	receptor.Handle()

	fmt.Println("\nVIP customer chain:")
	receptor.SetNext(manager)
	manager.SetNext(cashier)
	receptor.Handle()

}
