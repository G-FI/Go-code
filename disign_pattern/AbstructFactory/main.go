package main

import (
	"abstructFactory/abstructfactory"
	"fmt"
)

func appConfig() string {
	return "mac"
}
func getFactory() (abstructfactory.AbsFactory, error) {
	os := appConfig()
	if os == "win" {
		return abstructfactory.NewWinFactory(), nil
	} else if os == "mac" {
		return abstructfactory.NewMacFactory(), nil
	}
	return nil, fmt.Errorf("%s not support", os)
}

func main() {
	if f, err := getFactory(); err != nil {
		fmt.Println(err)
	} else {
		f.CreateButton().Press()
		f.CreateCheckBox().Check()
	}
}
