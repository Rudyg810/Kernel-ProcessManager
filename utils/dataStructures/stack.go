package dataStructures

import (
	"errors"
	"ProcessManger/types"
)

type Stack struct {
	items []types.Process
	top   int
}

func (s *Stack) push(item types.Process) {
	if s.items == nil {
		s.items = make([]types.Process, 0)
	}
	s.items = append(s.items, item)
	s.top = len(s.items) - 1
}
func (s *Stack) pop() (popped types.Process, err error) {
	if s.top == -1 {
		return types.Process{}, errors.New("stack is empty")
	}
	popped = s.items[s.top]
	s.items = s.items[:s.top]
	s.top = len(s.items) - 1
	return
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) peek() (peeked types.Process, err error){
	if s.top == -1{
		return types.Process{}, errors.New("stack is empty")
	}
	peeked = s.items[s.top]
	return
}
func (s *Stack) initialPop() (popped types.Process, err error) {
	if s.top == -1 {
		return types.Process{}, errors.New("stack is empty")
	}
	popped = s.items[0]
	s.items = s.items[1 : s.top+1]
	s.top = len(s.items) - 1
	return
}
func (s *Stack) display() []types.Process {
	return s.items
}

type StackImplementation interface {
	push(item types.Process)
	pop() (types.Process, error)
	isEmpty() bool
	peek() (types.Process, error)
	initialPop() (types.Process, error)
	display() []types.Process
}

