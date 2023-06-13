package factory

import "fmt"

const (
	ServerType = "server"
	PCType     = "pc"
	LaptopType = "labtop"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case PCType:
		return NewPC()
	case LaptopType:
		return NewLaptop()
	default:
		fmt.Printf("несуществующий тип объекта: %s", typeName)
		return nil
	}
}
