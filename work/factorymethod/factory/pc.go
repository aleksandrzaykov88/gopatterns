package factory

import "fmt"

type PC struct {
	Type    string
	Core    int
	Memory  int
	Display bool
}

func NewPC() Computer {
	return PC{
		Type:    PCType,
		Core:    8,
		Memory:  32,
		Display: true,
	}
}

func (s PC) GetType() string {
	return s.Type
}

func (s PC) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem [%d] Display %t", s.Type, s.Core, s.Memory, s.Display)
}
