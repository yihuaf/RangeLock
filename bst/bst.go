package bst

// Node ...
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert value into node.
func (n *Node) Insert(value int) error { return nil }

// Remove the value.
func (n *Node) Remove(value int) { return }

// Find the value in the BST. Error if not found.
func (n *Node) Find(value int) (*Node, error) { return nil, nil }

// IsBetween find if the value is in between two node inside the BST. In another word,
// can this value be inserted into one of the leaf node.
func (n *Node) IsBetween(value int) bool { return false }
