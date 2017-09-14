package rangelock_test

import (
	"RangeLock/rangelock"
	"testing"
)

func TestBasic(t *testing.T) {
	lock := rangelock.RangeLock{}
	if lock.Lock(3, 35) != nil {
		t.Error("Failed ..")
	}
}
