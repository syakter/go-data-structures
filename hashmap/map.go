package hashmap

import (
	g "github.com/syakter/go-data-structures/generic"
	"github.com/syakter/go-data-structures/iter"
)

type entry[K, V any] struct {
	key    K
	filled bool
	value  V
}

type Map[K, V any] struct {
	entries  []entry[K, V]
	capacity uint64
	length   uint64
	readonly bool
	ops      ops[K]
}

type ops[T any] struct {
	equals func(a, b T) bool
	hash   func(t T) uint64
}

func NewMap[K, V any](capacity uint64, equals g.EqualsFn[K], hash g.HashFn[K]) *Map[K, V] {
	if capacity == 0 {
		capacity = 1
	}
	return &Map[K, V]{
		entries:  make([]entry[K, V], capacity),
		capacity: capacity,
		ops: ops[K]{
			equals: equals,
			hash:   hash,
		},
	}
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	hash := m.ops.hash(key)
	idx := hash & (m.capacity - 1)

	for m.entries[idx].filled {
		if m.ops.equals(m.entries[idx].key, key) {
			return m.entries[idx].value, true
		}
		idx++
		if idx >= m.capacity {
			idx = 0
		}
	}

	var v V
	return v, false
}

func (m *Map[K, V]) resize(newcap uint64) {
	newm := Map[K, V]{
		capacity: newcap,
		length:   m.length,
		entries:  make([]entry[K, V], newcap),
		ops:      m.ops,
	}

	for _, ent := range m.entries {
		if ent.filled {
			newm.Put(ent.key, ent.value)
		}
	}
	m.capacity = newm.capacity
	m.entries = newm.entries
}

func (m *Map[K, V]) Put(key K, val V) {
	if m.length >= m.capacity/2 {
		m.resize(m.capacity * 2)
	} else if m.readonly {
		entries := make([]entry[K, V], len(m.entries))
		copy(entries, m.entries)
		m.entries = entries
	}

	hash := m.ops.hash(key)
	idx := hash & (m.capacity - 1)

	for m.entries[idx].filled {
		if m.ops.equals(m.entries[idx].key, key) {
			m.entries[idx].value = val
			return
		}
		idx++
		if idx >= m.capacity {
			idx = 0
		}
	}

	m.entries[idx].key = key
	m.entries[idx].value = val
	m.entries[idx].filled = true
	m.length++
}

func (m *Map[K, V]) remove(idx uint64) {
	var k K
	var v V
	m.entries[idx].filled = false
	m.entries[idx].key = k
	m.entries[idx].value = v
	m.length--
}

func (m *Map[K, V]) Remove(key K) {
	hash := m.ops.hash(key)
	idx := hash & (m.capacity - 1)

	for m.entries[idx].filled && !m.ops.equals(m.entries[idx].key, key) {
		idx = (idx + 1) % m.capacity
	}

	if !m.entries[idx].filled {
		return
	}

	m.remove(idx)

	idx = (idx + 1) % m.capacity
	for m.entries[idx].filled {
		krehash := m.entries[idx].key
		vrehash := m.entries[idx].value
		m.remove(idx)
		m.Put(krehash, vrehash)
		idx = (idx + 1) % m.capacity
	}

	if m.length > 0 && m.length <= m.capacity/8 {
		m.resize(m.capacity / 2)
	}
}

func (m *Map[K, V]) Size() int {
	return int(m.length)
}

func (m *Map[K, V]) Copy() *Map[K, V] {
	m.readonly = true
	return &Map[K, V]{
		entries:  m.entries,
		capacity: m.capacity,
		length:   m.length,
		readonly: m.readonly,
		ops:      m.ops,
	}
}

type KV[K, V any] struct {
	Key K
	Val V
}

func (m *Map[K, V]) Iter() iter.Iter[KV[K, V]] {
	kvs := make([]KV[K, V], 0, m.length)
	for _, ent := range m.entries {
		if ent.filled {
			kvs = append(kvs, KV[K, V]{
				Key: ent.key,
				Val: ent.value,
			})
		}
	}
	return iter.Slice(kvs)
}
