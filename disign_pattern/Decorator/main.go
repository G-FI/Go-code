package main

import (
	"Decorator/decorator"
	"fmt"
)

func main() {
	//original source
	source := decorator.NewDataSource()
	source.Write("xxxxxx")
	fmt.Println(source.Read())

	//Base Decorator
	source = decorator.NewDataSourceDecorator(source)
	fmt.Println(source.Read())

	//encryption Decorator
	source = decorator.NewEncryptionDecorator(source)
	source.Write("xxxxxe")
	fmt.Println(source.Read())

	//compress Decorator
	source = decorator.NewCompressDecorator(source)
	source.Write("xxxxxc")
	fmt.Println(source.Read())
}
