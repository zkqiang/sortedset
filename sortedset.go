package sortedset

import (
	"slices"
	"sort"
	"sync"
)

// SortedSet is a thread-safe sorted set implementation.
// It maintains both a map for O(1) lookups and a sorted slice for ordered iteration.
type SortedSet[T comparable] struct {
	elements map[T]struct{}
	order    []T
	sortFunc func(i, j T) bool
	mu       sync.RWMutex
}

func (s *SortedSet[T]) Add(element T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.elements[element]; !exists {
		s.elements[element] = struct{}{}
		index := sort.Search(len(s.order), func(i int) bool {
			return s.sortFunc(element, s.order[i])
		})
		s.order = slices.Insert(s.order, index, element)
	}
}

func (s *SortedSet[T]) AddAll(elements []T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, element := range elements {
		if _, exists := s.elements[element]; !exists {
			s.elements[element] = struct{}{}
			index := sort.Search(len(s.order), func(i int) bool {
				return s.sortFunc(element, s.order[i])
			})
			s.order = slices.Insert(s.order, index, element)
		}
	}
}

func (s *SortedSet[T]) Remove(element T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.elements[element]; exists {
		delete(s.elements, element)
		index := sort.Search(len(s.order), func(i int) bool {
			return !s.sortFunc(s.order[i], element) && !s.sortFunc(element, s.order[i])
		})
		s.order = append(s.order[:index], s.order[index+1:]...)
	}
}

func (s *SortedSet[T]) Contains(element T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.elements[element]
	return exists
}

func (s *SortedSet[T]) ContainsAll(elements []T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, element := range elements {
		if _, exists := s.elements[element]; !exists {
			return false
		}
	}
	return true
}

func (s *SortedSet[T]) ContainsAny(elements []T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, element := range elements {
		if _, exists := s.elements[element]; exists {
			return true
		}
	}
	return false
}

func (s *SortedSet[T]) Elements() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]T, len(s.order))
	copy(result, s.order)
	return result
}

func (s *SortedSet[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.order)
}

func (s *SortedSet[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.order) == 0
}

func (s *SortedSet[T]) PopLeft() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.order) == 0 {
		var zero T
		return zero, false
	}
	element := s.order[0]
	s.order = s.order[1:]
	delete(s.elements, element)
	return element, true
}

func (s *SortedSet[T]) PopRight() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.order) == 0 {
		var zero T
		return zero, false
	}
	element := s.order[len(s.order)-1]
	s.order = s.order[:len(s.order)-1]
	delete(s.elements, element)
	return element, true
}

func (s *SortedSet[T]) PopAll() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	elements := s.order
	s.order = []T{}
	s.elements = make(map[T]struct{})
	return elements
}

func (s *SortedSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.order = []T{}
	s.elements = make(map[T]struct{})
}

func (s *SortedSet[T]) Clone() *SortedSet[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	clone := &SortedSet[T]{
		elements: make(map[T]struct{}, len(s.elements)),
		order:    make([]T, len(s.order)),
		sortFunc: s.sortFunc,
	}
	copy(clone.order, s.order)
	for k := range s.elements {
		clone.elements[k] = struct{}{}
	}
	return clone
}
