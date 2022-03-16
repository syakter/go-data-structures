package cache

import (
	"github.com/syakter/go-data-structures/linkedlist"
)

type Cache[K comparable, V any] struct {
	size     int
	capacity int
	lru      linkedlist.LinkedList[KV[K, V]]
	table    map[K]*linkedlist.Node[KV[K, V]]
}

type KV[K comparable, V any] struct {
	Key K
	Val V
}

func New[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		size:     0,
		capacity: capacity,
		lru:      linkedlist.LinkedList[KV[K, V]]{},
		table:    make(map[K]*linkedlist.Node[KV[K, V]]),
	}
}
