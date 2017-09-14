package rangelock_test

import (
	"RangeLock/rangelock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLockBasic(t *testing.T) {
	lock := rangelock.New()
	assert.NoError(t, lock.Lock(3, 35))
	assert.NoError(t, lock.Lock(1, 2))
	assert.NoError(t, lock.Lock(44, 77))
	assert.Error(t, lock.Lock(44, 77))
	assert.Error(t, lock.Lock(45, 78))
	assert.Error(t, lock.Lock(22, 99))
}

func TestUnlock(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		lock := rangelock.New()
		assert.NoError(t, lock.Lock(3, 35))
		assert.NoError(t, lock.Lock(1, 2))
		assert.NoError(t, lock.Lock(44, 77))
		assert.NoError(t, lock.Unlock(3, 35))
		assert.NoError(t, lock.Unlock(1, 2))
		assert.NoError(t, lock.Unlock(44, 77))
	})

	t.Run("Shouldn't unlock", func(t *testing.T) {
		lock := rangelock.New()
		assert.NoError(t, lock.Lock(3, 35))
		assert.NoError(t, lock.Lock(1, 2))
		assert.NoError(t, lock.Lock(44, 77))
		assert.Error(t, lock.Unlock(2, 44))
		assert.Error(t, lock.Unlock(1, 1))
		assert.Error(t, lock.Unlock(33, 77))
		assert.Error(t, lock.Unlock(79, 800))
	})
}
