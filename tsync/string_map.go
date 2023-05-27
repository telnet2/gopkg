package tsync

import "sync"

// StringMap is a mutex guarded map.
type StringMap[T any] struct {
	m map[string]T
	l sync.RWMutex
}

func NewStringMap[T any]() *StringMap[T] {
	return &StringMap[T]{
		m: make(map[string]T),
	}
}

func (s *StringMap[T]) Set(key string, value T) {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[key] = value
}

func (s *StringMap[T]) Has(key string) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	_, ok := s.m[key]
	return ok
}

func (s *StringMap[T]) Get(key string) (T, bool) {
	s.l.RLock()
	defer s.l.RUnlock()
	v, ok := s.m[key]
	return v, ok
}

func (s *StringMap[T]) Remove(key string) {
	s.l.Lock()
	defer s.l.Unlock()
	delete(s.m, key)
}
