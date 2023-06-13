package composite

import (
	"fmt"
	"testing"
)

func TestCompositeTree(t *testing.T) {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
		},
		Left: &Tree{4, nil, nil},
	}

	fmt.Println(root.Right.Right.LeafValue)
}
