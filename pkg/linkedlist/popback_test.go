package linkedlist

import (
    "testing"
)

func TestPushOneValueAndPopBackOneValue(t *testing.T) {
	ll := New()
	ll.Push(3)
	ll.PopBack()
	size := ll.Size()
	if size != 0 {
		t.Errorf("wrong size, expected 0 and got %d", size)
	}
}

func TestPushTwoValueAndPopBackOneValue(t *testing.T) {
	ll := New()
	ll.Push(3)
	ll.Push(3)
	ll.PopBack()
	size := ll.Size()
	if size != 1 {
		t.Errorf("wrong size, expected 1 and got %d", size)
	}
}

func TestPushThreeValueAndPopBackOneValue(t *testing.T) {
	ll := New()
	ll.Push(3)
	ll.Push(3)
	ll.Push(3)
	ll.PopBack()
	size := ll.Size()
	if size != 2 {
		t.Errorf("wrong size, expected 2 and got %d", size)
	}
}

func TestPopBackTwoValueFromEmptyLinkedList(t *testing.T) {
	ll := New()
	ll.PopBack()
	ll.PopBack()
	size := ll.Size()
	if size != 0 {
		t.Errorf("wrong size, expected 0 and got %d", size)
	}
}