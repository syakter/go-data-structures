package hashset

import (
	g "github.com/syakter/go-data-structures/generic"
	"github.com/syakter/go-data-structures/hashmap"
	"github.com/syakter/go-data-structures/iter"
)

type Set[K any] struct {
	m *hashmap.Map[K, struct{}]
}

func New[K any](capacity uint64, equals g.EqualsFn[K], hash g.HashFn[K]) *Set[K] {
	return &Set[K]{
		m: hashmap.NewMap[K, struct{}](capacity, equals, hash),
	}
}

func (s *Set[K]) Put(val K) {
	s.m.Put(val, struct{}{})
}

func (s *Set[K]) Has(val K) bool {
	_, ok := s.m.Get(val)
	return ok
}

func (s *Set[K]) Remove(val K) {
	s.m.Remove(val)
}

func (s *Set[K]) Size() int {
	return s.m.Size()
}

func (s *Set[K]) Iter() iter.Iter[K] {
	it := s.m.Iter()
	return func() (K, bool) {
		kv, ok := it()
		return kv.Key, ok
	}
}
