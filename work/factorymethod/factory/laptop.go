package factory

import "fmt"

type Laptop struct {
	Type   string
	Core   int
	Memory int
}

func NewLaptop() Computer {
	return Laptop{
		Type:   LaptopType,
		Core:   8,
		Memory: 8,
	}
}

func (s Laptop) GetType() string {
	return s.Type
}

func (s Laptop) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem [%d]", s.Type, s.Core, s.Memory)
}
