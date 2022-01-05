package linked_list

import (
	"errors"
	"fmt"
)

type ele struct {
	name string
	next *ele
}

type singlelist struct {
	len  int
	head *ele
}

func CreateSignleList() *singlelist {
	return &singlelist{}
}

func (s *singlelist) AddFront(v string) {
	el := &ele{
		name: v,
	}

	if s.head == nil {
		s.head = el
	} else {
		el.next = s.head
		s.head = el
	}
	s.len++
}

func (s *singlelist) AddBack(v string) {
	el := &ele{
		name: v,
	}

	if s.head == nil {
		s.head = el
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = el
	}
	s.len++
}

func (s *singlelist) RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("This is an empty list")
	} else {
		var prev *ele
		current := s.head
		for current.next != nil {
			prev = current
			current = current.next
		}
		prev.next = nil
		return nil
	}
}

func (s *singlelist) Traverse() error {
	if s.head == nil {
		return errors.New("Empty linked list")
	} else {
		current := s.head
		for current != nil {
			fmt.Println(current.name)
			current = current.next
		}
		return nil
	}
}
