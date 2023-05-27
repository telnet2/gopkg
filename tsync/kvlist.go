package tsync

import (
	"fmt"
	"sort"
	"sync"
)

type KeyType interface {
	~int | ~int32 | ~int64 | ~string
}

type elem[K KeyType, V any] struct {
	Key K
	Val V
}

type SortedKV[K KeyType, V any] struct {
	l   sync.RWMutex
	arr []*elem[K, V]
}

// Store inserts a value with a given key if there's no value in the array.
// If a key is already in the list, then update the value.
func (s *SortedKV[K, V]) Store(key K, val V) {
	s.l.Lock()
	defer s.l.Unlock()

	idx, ok := sort.Find(len(s.arr), func(i int) int {
		if s.arr[i].Key == key {
			return 0
		} else if s.arr[i].Key < key {
			return 1
		} else {
			return -1
		}
	})
	if ok {
		// update the value because key is the same
		s.arr[idx].Val = val
	} else {
		if idx == len(s.arr) {
			s.arr = append(s.arr, &elem[K, V]{Key: key, Val: val})
		} else {
			s.arr = append(s.arr[:idx+1], s.arr[idx:]...)
			s.arr[idx] = &elem[K, V]{Key: key, Val: val}
		}
	}
}

// Load returns a value for a given key.
// If it cannot find a value for the key, it returns false.
func (s *SortedKV[K, V]) Load(key K) (ret V, ok bool) {
	s.l.RLock()
	defer s.l.RUnlock()

	idx := sort.Search(len(s.arr), func(i int) bool {
		return s.arr[i].Key >= key
	})
	if idx < 0 || idx == len(s.arr) {
		return ret, false
	}
	if s.arr[idx].Key != key {
		return ret, false
	}
	return s.arr[idx].Val, true
}

// RangePart calls the visit function for the interval given page and limit.
// page is 1-based number.
func (s *SortedKV[K, V]) RangePart(page, limit uint32, visit func(k K, v V) bool) {
	if page < 1 {
		page = 1
	}
	page -= 1
	for i := page * limit; i < (page+1)*limit; i++ {
		s.l.RLock()
		if int(i) < len(s.arr) {
			if !visit(s.arr[i].Key, s.arr[i].Val) {
				s.l.RUnlock()
				break
			}
		}
		s.l.RUnlock()
	}
}

// RangePart calls the visit function for the interval given page and limit.
// page is 1-based number.
func (s *SortedKV[K, V]) Range(visit func(k K, v V) bool) {
	i := 0
	for {
		s.l.RLock()
		if i < len(s.arr) {
			if !visit(s.arr[i].Key, s.arr[i].Val) {
				s.l.RUnlock()
				break
			}
			i++
		} else {
			s.l.RUnlock()
			break
		}
		s.l.RUnlock()
	}
}

// Keys returns an array of keys as a string
func (s *SortedKV[K, V]) Keys() []string {
	s.l.RLock()
	defer s.l.RUnlock()
	var ret []string
	for _, e := range s.arr {
		ret = append(ret, fmt.Sprintf("%v", e.Key))
	}
	return ret
}

// Size returns the size of the map.
func (s *SortedKV[K, V]) Size() int {
	s.l.RLock()
	defer s.l.RUnlock()
	return len(s.arr)
}
