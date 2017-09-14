package rangelock

import (
	"RangeLock/bst"
	"errors"
	"sync"
)

// RangeLock ....
type RangeLock struct {
	mutex *sync.Mutex
	bst   *bst.BST
}

// New ...
func New() *RangeLock {
	return &RangeLock{
		mutex: &sync.Mutex{},
		bst:   bst.New(),
	}
}

// Lock ...
func (rl *RangeLock) Lock(begin, end int) error {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	// Search for the range in bst.
	if err := rl.bst.Insert(begin, end); err != nil {
		return errors.New("Locked already")
	}

	return nil
}

// Unlock ...
func (rl *RangeLock) Unlock(begin, end int) error {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	node, err := rl.bst.Find(begin, end)
	if err != nil {
		return errors.New("Can only unlock exact locked range")
	}

	if node.Begin != begin || node.End != end {
		return errors.New("Can only unlock exact locked range")
	}

	rl.bst.Remove(begin, end)
	return nil
}
