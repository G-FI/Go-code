package main

import (
	"factory/factory"
	"fmt"
)

//客户端代码
func main() {
	gun1, _ := factory.CreateGun("ak47")
	gun2, _ := factory.CreateGun("m416")

	showGun(gun1)
	showGun(gun2)
}

func showGun(gun factory.IGun) {
	fmt.Println(gun.GetName())
	fmt.Println(gun.GetPower())
}
