package bst

// Node ...
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	tree *Node
}

func New() *BST {
	return &BST{}
}

// Insert value into node.
func (n *BST) Insert(value int) error { return nil }

// Remove the value.
func (n *BST) Remove(value int) { return }

// Find the value in the BST. Error if not found.
func (n *BST) Find(value int) (*Node, error) { return nil, nil }

// IsBetween find if the value is in between two node inside the BST. In another word,
// can this value be inserted into one of the leaf node.
func (n *BST) IsBetween(value int) bool { return false }
