package rangelock_test

import (
	"RangeLock/rangelock"
	"testing"
)

func TestBasic(t *testing.T) {
	lock := rangelock.New()
	if err := lock.Lock(3, 35); err != nil {
		t.Fatal("Should be able to lock", err)
	}

	if err := lock.Lock(4, 5); err == nil {
		t.Fatal("Should not lock")
	}

	if err := lock.Unlock(3, 35); err != nil {
		t.Fatal("Should unlock")
	}
}
