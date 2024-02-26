package DSA_Go

import (
	"testing"
)

func assertEqual[T comparable](t *testing.T, expected T, actual T) {
	t.Helper()
	if expected == actual {
		return
	}
	t.Errorf("expected (%+v) is not equal to actual (%+v)", expected, actual)
}

func TestLinkedList(t *testing.T) {
	head := &LinkedList{}
	assertEqual(t, 0, head.length())

	v1, e1 := head.get(0)
	if e1 == nil {
		t.Errorf("expected error for empty list")
	}
	assertEqual(t, 0, v1)

	v2, e2 := head.get(1)
	if e2 == nil {
		t.Errorf("expected error for empty list")
	}
	assertEqual(t, 0, v2)

	head.append(1)
	v3, e3 := head.get(0)
	if e3 != nil {
		t.Errorf("expected valid list to not error")
	}
	assertEqual(t, 1, v3)
	assertEqual(t, 1, head.length())

	head.append(2)
	head.append(3)
	head.append(4)

	head.prepend(0)
	head.prepend(-1)

	v4, e4 := head.get(0)
	if e4 != nil {
		t.Errorf("expected valid list to not error")
	}
	assertEqual(t, -1, v4)

	v5, e5 := head.pop()
	if e5 != nil {
		t.Errorf("expected valid list to not error")
	}
	assertEqual(t, 4, v5)
	head.print()

	v6, e6 := head.remove(0)
	if e6 != nil {
		t.Errorf(e6.Error())
	}
	assertEqual(t, -1, v6)

	v7, e7 := head.remove(1)
	if e7 != nil {
		t.Errorf(e7.Error())
	}
	assertEqual(t, 1, v7)

	v8, e8 := head.remove(4)
	if e8 == nil {
		t.Errorf("Error should not be nil")
	}
	assertEqual(t, 0, v8)
}
