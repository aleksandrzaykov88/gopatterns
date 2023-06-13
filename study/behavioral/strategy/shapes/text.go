package shapes

import (
	"design_patterns/behavioral/strategy"
)

type TextSquare struct {
	strategy.PrintOutput
}

func (t *TextSquare) Print() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
