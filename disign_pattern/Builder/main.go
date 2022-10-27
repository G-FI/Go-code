package main

import "Builder/builder"

func main() {
	manager := builder.NewManager()

	//build a car
	b, _ := builder.NewBuilder("car")
	manager.Make(b)
	carBuilder := b.(*builder.CarBuilder)
	car := carBuilder.GetProduct()
	car.Run()

	//build a car manual
	b, _ = builder.NewBuilder("carmanual")
	manager.Make(b)
	carManualBuilder := b.(*builder.CarManualBuilder)
	carManual := carManualBuilder.GetProduct()
	carManual.ShowDetile()
}
