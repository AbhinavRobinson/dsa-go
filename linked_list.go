package DSA_Go

import "errors"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) append(value int) {
	newNode := &Node{data: value}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) length() int {
	length := 0
	if list.head == nil {
		return length
	} else {
		length++
		current := list.head
		for current.next != nil {
			length++
		}
	}
	return length
}

func (list *LinkedList) get(index int) (int, error) {
	current := list.head
	if current == nil {
		return 0, errors.New("empty list")
	}
	for current.next != nil {
		if index == 0 {
			return current.data, nil
		}
		index--
	}
	return current.data, nil
}
