package shapes

import (
	"design_patterns/behavioral/strategy"
	"fmt"
	"os"
)

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

func NewPrinter(s string) (strategy.PrintStrategy, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare{
			PrintOutput: strategy.PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			PrintOutput: strategy.PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("strategy '%s' not found", s)
	}
}
