package DSA_Go

import "testing"

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
}
