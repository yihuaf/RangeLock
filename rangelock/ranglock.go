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
	if rl.bst.IsBetween(begin) || rl.bst.IsBetween(end) {
		return errors.New("Locked")
	}

	rl.bst.Insert(begin)
	rl.bst.Insert(end)
	return nil
}

// Unlock ...
func (rl *RangeLock) Unlock(begin, end int) error {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	_, foundBegin := rl.bst.Find(begin)
	_, foundEnd := rl.bst.Find(end)
	if foundBegin != nil && foundEnd != nil {
		rl.bst.Remove(begin)
		rl.bst.Remove(end)
		return nil
	}

	return errors.New("Didn't lock the exact range")
}
