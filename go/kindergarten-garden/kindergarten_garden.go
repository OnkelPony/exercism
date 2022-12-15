package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Garden represents places for plants and children, who take care of them.
type Garden struct {
	diagram  []rune
	children []string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

const cupsPerChild = 4

var validCups = map[rune]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

// NewGarden creates a new garden from the given diagram and children slice.
func NewGarden(diagram string, children []string) (*Garden, error) {
	if diagram[0] != '\n' || diagram[len(diagram)/2] != '\n' {
		return nil, errors.New("wrong diagram format")
	}

	diagRunes := []rune(strings.ReplaceAll(diagram, "\n", ""))

	if len(diagRunes)%4 != 0 {
		return nil, errors.New("wrong number of cups")
	}

	for _, cup := range diagRunes {
		if _, ok := validCups[cup]; !ok {
			return nil, errors.New("invalid cup code")
		}
	}

	partChildren := make(map[string]bool)
	for _, child := range children {
		if partChildren[child] {
			return nil, errors.New("duplicate name")
		}
		partChildren[child] = true
	}

	result := &Garden{
		diagram:  diagRunes,
		children: children,
	}
	return result, nil
}

// Plants method returns slice of plant names, belonging to provided child.
func (g *Garden) Plants(child string) ([]string, bool) {
	result := make([]string, cupsPerChild)

	sort.Strings(g.children)
	childIdx := make(map[string]int)
	for i, v := range g.children {
		childIdx[v] = i
	}

	square, ok := childIdx[child]
	if !ok {
		return result, false
	}

	result[0] = validCups[g.diagram[square*2]]
	result[1] = validCups[g.diagram[square*2+1]]
	result[2] = validCups[g.diagram[len(g.diagram)/2+square*2]]
	result[3] = validCups[g.diagram[len(g.diagram)/2+square*2+1]]

	return result, true
}
