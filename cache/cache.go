package cache

import (
	"github.com/syakter/go-data-structures/iter"
	"github.com/syakter/go-data-structures/list"
)

type Cache[K comparable, V any] struct {
	size     int
	capacity int
	lru      list.List[KV[K, V]]
	table    map[K]*list.Node[KV[K, V]]
}

type KV[K comparable, V any] struct {
	Key K
	Val V
}

func New[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		size:     0,
		capacity: capacity,
		lru:      list.List[KV[K, V]]{},
		table:    make(map[K]*list.Node[KV[K, V]]),
	}
}

func (t *Cache[K, V]) Get(k K) (V, bool) {
	if n, ok := t.table[k]; ok {
		t.lru.Remove(n)
		t.lru.PushFrontNode(n)
		return n.Value.Val, true
	}
	var v V
	return v, false
}

func (t *Cache[K, V]) Put(k K, e V) {
	if n, ok := t.table[k]; ok {
		n.Value.Val = e
		t.lru.Remove(n)
		t.lru.PushFrontNode(n)
		return
	}

	if t.size == t.capacity {
		t.evict()
	}
	n := &list.Node[KV[K, V]]{
		Value: KV[K, V]{
			Key: k,
			Val: e,
		},
	}
	t.lru.PushFrontNode(n)
	t.size++
	t.table[k] = n
}

func (t *Cache[K, V]) evict() {
	key := t.lru.Back.Value.Key
	t.lru.Remove(t.lru.Back)
	t.size--
	delete(t.table, key)
}

func (t *Cache[K, V]) Remove(k K) {
	if n, ok := t.table[k]; ok {
		t.lru.Remove(n)
		t.size--
		delete(t.table, k)
	}
}

func (t *Cache[K, V]) Resize(size int) {
	if t.capacity == size {
		return
	} else if t.capacity < size {
		t.capacity = size
	}

	for i := 0; i < t.capacity-size; i++ {
		t.evict()
	}

	t.capacity = size
}

func (t *Cache[K, V]) Size() int {
	return t.size
}

func (t *Cache[K, V]) Capacity() int {
	return t.capacity
}

func (t *Cache[K, V]) Iter() iter.Iter[KV[K, V]] {
	return t.lru.Front.Iter()
}
