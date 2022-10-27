package main

import (
	"composite/composite"
	"fmt"
)

func main() {
	circle := composite.NewCircle(0, 0, 10)
	rectangle := composite.NewRectangle(10, 10, 100, 5)
	circle.Move(1, 1)

	circle.Draw()
	rectangle.Draw()
	fmt.Println("------------------------------------------")

	compoundGraphic := composite.NewCompoundGraphic()
	compoundGraphic.Add(circle)
	compoundGraphic.Add(rectangle)

	compoundGraphic.Move(-100, -100)
	compoundGraphic.Draw()
	fmt.Println("------------------------------------------")

	compoundGraphic.Scale(10)
	compoundGraphic.Draw()

	fmt.Println("-------------------------------------------")

	compoundGraphic.Scale(0.05)
	compoundGraphic.Draw()
}
