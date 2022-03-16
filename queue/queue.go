package queue

import (
	"github.com/syakter/go-data-structures/iter"
	"github.com/syakter/go-data-structures/list"
)

type Queue[T any] struct {
	list *list.List[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		list: list.New[T](),
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.PushBack(value)
}

func (q *Queue[T]) Dequeue() T {
	value := q.list.Front.Value
	q.list.Remove(q.list.Front)

	return value
}

func (q *Queue[T]) Peek() T {
	return q.list.Front.Value
}

func (q *Queue[T]) Empty() bool {
	return q.list.Front == nil
}

func (q *Queue[T]) Iter() iter.Iter[T] {
	return q.list.Front.Iter()
}
