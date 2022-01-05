package set

import "fmt"

type Set struct {
	container map[string]struct{}
}

func MakeSet() *Set {
	return &Set{
		container: make(map[string]struct{}),
	}
}

func (s *Set) Add(v string) {
	if _, ok := s.container[v]; !ok {
		s.container[v] = struct{}{}
	}
}

func (s *Set) Remove(v string) error {
	if _, ok := s.container[v]; ok {
		delete(s.container, v)
		return nil
	}
	return fmt.Errorf("Error while deleting element...")
}

func (s *Set) Exists(v string) bool {
	_, ok := s.container[v]
	return ok
}
