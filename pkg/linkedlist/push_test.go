package linkedlist

import (
    "testing"
)

func TestPushOneValue(t *testing.T) {
	ll := New()
	ll.Push(7)
}

func TestPushTwoValue(t *testing.T) {
	ll := New()
	ll.Push(9)
	ll.Push(5)
}

func TestPushThreeValue(t *testing.T) {
	ll := New()
	ll.Push(5)
	ll.Push(100)
	ll.Push(57)
}

