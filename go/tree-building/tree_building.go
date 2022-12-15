// Package tree implements functions to translate records of ID and Parent to the tree structure.
package tree

import (
	"errors"
	"sort"
)

// Record represents one ID and Parent pair, imported from database.
type Record struct {
	ID     int
	Parent int
}

// Node represents one node of a tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build makes Node structure from the slice of Record.
func Build(input []Record) (*Node, error) {
	var nodes = map[int]*Node{}
	sort.Slice(input, func(i, j int) bool {
		return input[i].ID < input[j].ID
	})

	for i, rec := range input {
		if rec.ID != i || rec.Parent > rec.ID || rec.ID > 0 && rec.Parent == rec.ID {
			return nil, errors.New("not in sequence or has bad parent")
		}

		nodes[rec.ID] = &Node{ID: rec.ID}

		if rec.ID > 0 {
			parent := nodes[rec.Parent]
			parent.Children = append(parent.Children, nodes[rec.ID])
		}

	}

	return nodes[0], nil
}
