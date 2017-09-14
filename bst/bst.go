// Package bst contains a binary search tree whose value is a range with a begin and an end.
package bst

import "errors"

// Node ...
type Node struct {
	Begin int
	End   int
	Left  *Node
	Right *Node
}

// IsOverlap test if any range between begin and end overlap with the node.
// Assume begin >= end
func (n *Node) IsOverlap(begin, end int) bool {
	return begin >= n.Begin || end < n.End || (begin <= n.Begin && end >= n.End)
}

// BST ...
type BST struct {
	tree *Node
}

// New ...
func New() *BST {
	return &BST{}
}

// Insert value into node.
func (bst *BST) Insert(begin, end int) error {
	if begin > end {
		return errors.New("Begin must be smaller than end")
	}

	node := &Node{Begin: begin, End: end}
	if bst.tree == nil {
		bst.tree = node
		return nil
	}

	for cur := bst.tree; cur != nil; {
		// If the range is completely to the left
		if end < cur.Begin {
			if cur.Left == nil {
				cur.Left = node
				return nil
			}

			cur = cur.Left
			continue
		}

		// If the range is completely to the right
		if begin > cur.End {
			if cur.Right == nil {
				cur.Right = node
				return nil
			}

			cur = cur.Right
			continue
		}

		return errors.New("Cannot Insert")
	}

	return nil
}

// Find the min value of the right subtree.
func findMin(subtree *Node) *Node {
	if subtree.Left == nil {
		return subtree
	}

	return findMin(subtree.Right)
}

func del(subtree *Node, begin, end int) *Node {
	if subtree == nil {
		return nil
	}

	if end < subtree.Begin {
		subtree.Left = del(subtree.Left, begin, end)
		return subtree
	}

	if begin > subtree.End {
		subtree.Right = del(subtree.Right, begin, end)
		return subtree
	}

	// Will only delete the node if begin and end is exact match to
	// the node.
	if begin == subtree.Begin && end == subtree.End {
		// Leaf node
		if subtree.Left == nil && subtree.Right == nil {
			return nil
		}

		// One subtree exist
		if subtree.Left == nil {
			return subtree.Right
		}

		if subtree.Right == nil {
			return subtree.Left
		}

		// Both left and right subtree are not nil
		minNode := findMin(subtree)
		minNode.Right = del(subtree.Right, minNode.Begin, minNode.End)
		minNode.Left = subtree.Left
		return minNode
	}

	return subtree
}

// Remove the value. Will only remove if begin and end is exact match.
func (bst *BST) Remove(begin, end int) {
	if bst.tree == nil {
		return
	}

	bst.tree = del(bst.tree, begin, end)
}

// Find a node thats overlap with the given range in the BST. Error if not found.
func (bst *BST) Find(begin, end int) (*Node, error) {
	if begin > end {
		return nil, errors.New("Begin must be smaller than end")
	}

	if bst.tree == nil {
		return nil, errors.New("Not Found")
	}

	for cur := bst.tree; cur != nil; {
		// If the range is completely to the left
		if end < cur.Begin {
			cur = cur.Left
			continue
		}

		// If the range is completely to the right
		if begin > cur.End {
			cur = cur.Right
			continue
		}

		return cur, nil
	}

	return nil, errors.New("Not Found")
}

func size(subtree *Node) int {
	if subtree == nil {
		return 0
	}

	return size(subtree.Left) + size(subtree.Right) + 1
}

// Size calculate the size of the tree.
func (bst *BST) Size() int {
	return size(bst.tree)
}
