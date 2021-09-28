package goexpression

import (
	"fmt"
	"sync"
)

// Item the type of the stack
type Item interface{}

// itemStack the stack of Items
type itemStack struct {
	items []Item
	lock  sync.RWMutex
}

// NewStack creates a new itemStack
func NewStack() *itemStack {
	s := &itemStack{}
	s.items = []Item{}
	return s
}

// Print prints all the elements
func (s *itemStack) Print() {
	fmt.Println(s.items)
}

// Push adds an Item to the top of the stack
func (s *itemStack) Push(t Item) {
	s.lock.Lock()
	s.lock.Unlock()
	s.items = append(s.items, t)
}

// Pop removes an Item from the top of the stack
func (s *itemStack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

//Peek get top Item
func (s *itemStack) Peek() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.items[len(s.items)-1]
}

//Size get size
func (s *itemStack) Size() int {
	return len(s.items)
}

//IsEmpty true if stack is empty, else false
func (s *itemStack) IsEmpty() bool {
	return len(s.items) == 0
}
