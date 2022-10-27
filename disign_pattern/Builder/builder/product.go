package builder

import "fmt"

type Car struct {
	seats  int
	engine string
}

func (c *Car) Run() {
	fmt.Printf("Car is runing with %d seats and %s engine\n", c.seats, c.engine)
}

type CarManual struct {
	seats  int
	engine string
}

func (c *CarManual) ShowDetile() {
	fmt.Printf("Car manual: %d seats, %s engine", c.seats, c.engine)
}
