package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	stack *list.List
}

func (s *Stack) Push(v string) {
	s.stack.PushFront(v)
}

func (s *Stack) Pop() error {
	if s.stack.Len() > 0 {
		el := s.stack.Front()
		s.stack.Remove(el)
	}

	return fmt.Errorf("Stack is empty...")
}

func (s *Stack) Front() (string, error) {
	if s.stack.Len() > 0 {
		if val, ok := s.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("Shit happens...")
	}

	return "", fmt.Errorf("The stack is empty...")
}

func (s *Stack) Size() int {
	return s.stack.Len()
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}
