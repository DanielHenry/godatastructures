package linkedlist

import (
    "testing"
)

func TestSizeAfterPushOneValue(t *testing.T) {
	ll := New()
	ll.Push(3)
	size := ll.Size()
	if size != 1 {
		t.Errorf("wrong size, expected 1 and got %d", size)
	}
}

func TestSizeAfterPushTwoValue(t *testing.T) {
	ll := New()
	ll.Push(3)
	size := ll.Size()
	if size != 1 {
		t.Errorf("wrong size, expected 1 and got %d", size)
	}
	ll.Push(3)
	size = ll.Size()
	if size != 2 {
		t.Errorf("wrong size, expected 2 and got %d", size)
	}
}

func TestSizeAfterPushThreeValue(t *testing.T) {
	ll := New()
	ll.Push(3)
	size := ll.Size()
	if size != 1 {
		t.Errorf("wrong size, expected 1 and got %d", size)
	}
	ll.Push(3)
	size = ll.Size()
	if size != 2 {
		t.Errorf("wrong size, expected 2 and got %d", size)
	}
	ll.Push(3)
	size = ll.Size()
	if size != 3 {
		t.Errorf("wrong size, expected 3 and got %d", size)
	}
}

