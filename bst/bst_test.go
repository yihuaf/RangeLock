package bst_test

import (
	"RangeLock/bst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertBasic(t *testing.T) {
	tree := bst.New()
	assert.NoError(t, tree.Insert(3, 55))
	assert.Error(t, tree.Insert(4, 5))
	assert.Error(t, tree.Insert(1, 4))
	assert.Error(t, tree.Insert(53, 70))
	assert.NoError(t, tree.Insert(1, 2))
	assert.NoError(t, tree.Insert(57, 59))
	assert.Error(t, tree.Insert(66, 65))
	return
}

func TestSize(t *testing.T) {
	tree := bst.New()
	assert.Equal(t, 0, tree.Size())
	assert.NoError(t, tree.Insert(3, 55))
	assert.Equal(t, 1, tree.Size())
	assert.NoError(t, tree.Insert(1, 2))
	assert.Equal(t, 2, tree.Size())
	assert.NoError(t, tree.Insert(222, 555))
	assert.Equal(t, 3, tree.Size())
}

func TestInsertLarge(t *testing.T) {
	tree := bst.New()
	assert.NoError(t, tree.Insert(30, 55))
	assert.NoError(t, tree.Insert(22, 25))
	assert.NoError(t, tree.Insert(1, 2))
	assert.NoError(t, tree.Insert(57, 59))
	assert.NoError(t, tree.Insert(60, 62))
	assert.NoError(t, tree.Insert(128, 999))
	assert.NoError(t, tree.Insert(77, 99))
	return
}
func TestFind(t *testing.T) {
	tree := bst.New()
	assert.NoError(t, tree.Insert(3, 55))
	var err error
	_, err = tree.Find(44, 22)
	assert.Error(t, err)
	_, err = tree.Find(1, 2)
	assert.Error(t, err)
	_, err = tree.Find(555, 666)
	assert.Error(t, err)

	var node *bst.Node

	node, err = tree.Find(3, 55)
	assert.NoError(t, err)
	assert.Equal(t, 3, node.Begin)
	assert.Equal(t, 55, node.End)

	node, err = tree.Find(4, 5)
	assert.NoError(t, err)
	assert.Equal(t, 3, node.Begin)
	assert.Equal(t, 55, node.End)
}

func TestOverlap(t *testing.T) {
	node := &bst.Node{Begin: 66, End: 77}
	assert.True(t, node.IsOverlap(33, 67))
	assert.True(t, node.IsOverlap(68, 99))
	assert.True(t, node.IsOverlap(1, 99))
}

func TestRemove(t *testing.T) {
	t.Run("Delete Basic", func(t *testing.T) {
		tree := bst.New()
		assert.NoError(t, tree.Insert(3, 55))
		assert.NoError(t, tree.Insert(1, 2))
		assert.NoError(t, tree.Insert(57, 59))
		assert.Equal(t, 3, tree.Size())
		tree.Remove(3, 55)
		assert.Equal(t, 2, tree.Size())
		tree.Remove(1, 2)
		assert.Equal(t, 1, tree.Size())
		tree.Remove(57, 59)
		assert.Equal(t, 0, tree.Size())

	})

	t.Run("Remove Left", func(t *testing.T) {
		tree := bst.New()
		assert.NoError(t, tree.Insert(3, 55))
		assert.NoError(t, tree.Insert(1, 2))
		assert.NoError(t, tree.Insert(57, 59))
		assert.Equal(t, 3, tree.Size())
		tree.Remove(1, 2)
		assert.Equal(t, 2, tree.Size())
		tree.Remove(3, 55)
		assert.Equal(t, 1, tree.Size())
		tree.Remove(57, 59)
		assert.Equal(t, 0, tree.Size())
	})

	t.Run("Remove Right", func(t *testing.T) {
		tree := bst.New()
		assert.NoError(t, tree.Insert(3, 55))
		assert.NoError(t, tree.Insert(1, 2))
		assert.NoError(t, tree.Insert(57, 59))
		assert.Equal(t, 3, tree.Size())
		tree.Remove(57, 59)
		assert.Equal(t, 2, tree.Size())
		tree.Remove(3, 55)
		assert.Equal(t, 1, tree.Size())
		tree.Remove(1, 2)
		assert.Equal(t, 0, tree.Size())
	})

	t.Run("Remove Not Found", func(t *testing.T) {
		tree := bst.New()
		assert.NoError(t, tree.Insert(3, 55))
		assert.NoError(t, tree.Insert(1, 2))
		assert.NoError(t, tree.Insert(57, 59))
		assert.Equal(t, 3, tree.Size())
		tree.Remove(4, 7)
		assert.Equal(t, 3, tree.Size())
		tree.Remove(3, 99)
		assert.Equal(t, 3, tree.Size())
		tree.Remove(1, 44)
		assert.Equal(t, 3, tree.Size())
	})
}
