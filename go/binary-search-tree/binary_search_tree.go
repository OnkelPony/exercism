// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package binarysearchtree implements structures and methods for operating on binary search tree.
package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// NewBst creates and returns a new SearchTreeData.
func NewBst(i int) SearchTreeData {
	return SearchTreeData{
		data: i,
	}
}

// Insert inserts an int into the SearchTreeData.
// Inserts happen based on the rules of a BinarySearchTree
func (std *SearchTreeData) Insert(i int) {
	cur := NewBst(i)
	switch {
	case i > std.data:
		if std.right == nil {
			std.right = &cur
		} else {
			std.right.Insert(i)
		}
	case i <= std.data:
		if std.left == nil {
			std.left = &cur
		} else {
			std.left.Insert(i)
		}
	}
}

// MapString returns the ordered contents of SearchTreeData as a []string.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []string ["1", "3", "5", "7"].
func (std *SearchTreeData) MapString(f func(int) string) (result []string) {
	var takeNext func(*SearchTreeData)
	takeNext = func(node *SearchTreeData) {
		if node.left != nil {
			takeNext(node.left)
		}
		result = append(result, f(node.data))
		if node.right != nil {
			takeNext(node.right)
		}
	}
	takeNext(std)
	return result
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(f func(int) int) (result []int) {
	var takeNext func(*SearchTreeData)
	takeNext = func(node *SearchTreeData) {
		if node.left != nil {
			takeNext(node.left)
		}
		result = append(result, f(node.data))
		if node.right != nil {
			takeNext(node.right)
		}
	}
	takeNext(std)
	return result
}
