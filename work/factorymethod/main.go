package main

import (
	"factorymethod/factory"
)

func main() {
	myPc := factory.New("pc")
	myPc.PrintDetails()

	// var _ factory.Computer = (*factory.Laptop)(nil)
}
