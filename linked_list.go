package DSA_Go

import (
	"errors"
	"fmt"
	"log"
)

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
			current = current.next
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
		current = current.next
		index--
	}
	if index == 0 {
		return current.data, nil
	} else {
		return 0, errors.New("out of range")
	}
}

func (list *LinkedList) prepend(value int) {
	current := list.head
	list.head = &Node{
		data: value,
		next: current,
	}
}

func (list *LinkedList) remove(index int) (int, error) {
	current := list.head
	if index > list.length()-1 {
		return 0, errors.New("index out of range")
	}
	if index == 0 {
		list.head = list.head.next
		return current.data, nil
	}
	if current == nil {
		return 0, errors.New("empty list")
	}
	if current.next == nil {
		list.head = nil
		return current.data, nil
	}
	prev := current
	for index != 0 {
		prev = current
		current = current.next
		index--
	}
	prev.next = current.next
	current.next = nil
	return current.data, nil
}

func (list *LinkedList) pop() (int, error) {
	current := list.head
	var prev *Node
	for current.next != nil {
		prev = current
		current = current.next
	}
	prev.next = nil
	return current.data, nil
}

func (list *LinkedList) print() {
	p := list.head
	v := ""
	for p != nil {
		v += fmt.Sprintf("-> %v ", p.data)
		p = p.next
	}
	log.Println(v)
}
