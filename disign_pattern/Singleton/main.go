package main

import (
	"fmt"
	"singleton/singleton"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			instance := singleton.GetInstance()
			instance.DoSomething()
		}()
	}
	fmt.Scanln()
}
