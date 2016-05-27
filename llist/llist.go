package llist

import (
	"strconv"
	"strings"
)

type node struct {
	val  int
	next *node
}

func (n *node) Next() *node {
	return n.next
}

// LList is a Linked List object
type LList struct {
	head *node
	size int
}

// New creates a new Linked List
func New() *LList {
	var l LList
	l.size = 0
	return &l
}

// Add adds an element towards the end of the linked list
func (l *LList) Add(num int) {
	n := &node{val: num}
	l.size++
	if l.head == nil {
		l.head = n

	} else {

		curr := l.head

		for curr.next != nil {
			curr = curr.Next()
		}

		curr.next = n

	}

}

// Delete deletes the specific node containing such element instead of using position for delete
func (l *LList) Delete(num int) bool {
	found := false
	curr := l.head
	if curr.val == num {
		l.head = curr.next
		found = true
		l.size--
	}
	for curr.next != nil {
		if curr.next.val == num {
			prev := curr
			curr = curr.next
			prev.next = curr.next
			curr = nil
			found = true
			l.size--
			break
		}
		curr = curr.next
	}
	return found
}

// Contains check to see if the element is in the Linked List
func (l *LList) Contains(num int) bool {
	found := false
	curr := l.head
	if curr.val == num {
		found = true
	}
	for curr.next != nil {
		if curr.next.val == num {
			found = true
			break
		}
		curr = curr.next
	}
	return found
}

// Size returns how big the Linked List is
func (l *LList) Size() int {
	return l.size
}

// String prints all the elements in the Linked List
func (l *LList) String() string {
	str := ""
	curr := l.head

	for curr != nil {
		str += strconv.Itoa(curr.val) + " "
		curr = curr.next
	}
	return strings.TrimRight(str, " ")
}
