package tsync

import (
	"container/list"
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedKV(t *testing.T) {
	n := make([]int, 10)
	l := new(SortedKV[int, int])

	for i := 0; i < len(n); i++ {
		n[i] = i
	}
	rand.Shuffle(len(n), func(i, j int) {
		n[i], n[j] = n[j], n[i]
	})

	for i := 0; i < len(n); i++ {
		l.Store(n[i], n[i])
	}

	_, ok := l.Load(-1)
	require.False(t, ok)

	_, ok = l.Load(1e3)
	require.False(t, ok)

	for i := 0; i < len(n); i++ {
		v, ok := l.Load(n[i])
		require.True(t, ok)
		require.Equal(t, n[i], v)
	}

	require.Equal(t, []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, l.Keys())

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			var part []int
			l.RangePart(3, 3, func(k, v int) bool {
				// 6, 7, 8
				part = append(part, v)
				return true
			})
			require.Equal(t, []int{6, 7, 8}, part)

			var all []int
			l.Range(func(k, v int) bool {
				// 6, 7, 8
				all = append(all, v)
				return true
			})
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestSortedItem(t *testing.T) {
	type item struct {
		key int
		val int
		ts  int
		e   *list.Element
	}

	items := map[int]*item{}
	list := new(list.List)
	for i := 0; i < 10; i++ {
		items[i] = &item{key: i, val: i + 1, ts: i}
		items[i].e = list.PushBack(items[i])
	}

	// the goal is to reorder items by ts
	items[5].ts = 11
	list.Remove(items[5].e)
	items[5].e = list.PushBack(items[5])
}
