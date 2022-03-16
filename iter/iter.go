package iter

type Iter[T any] func() (T, bool)

func (it Iter[T]) For(fn func(t T)) {
	for val, ok := it(); ok; val, ok = it() {
		fn(val)
	}
}

func (it Iter[T]) ForBreak(fn func(t T) bool) {
	for val, ok := it(); ok; val, ok = it() {
		if !fn(val) {
			break
		}
	}
}

func Slice[T any](slice []T) Iter[T] {
	var i int
	return func() (t T, ok bool) {
		if i >= len(slice) {
			return t, false
		}

		r := slice[i]
		i++
		return r, true
	}
}

func SliceReverse[T any](slice []T) Iter[T] {
	i := len(slice) - 1
	return func() (t T, ok bool) {
		if i < 0 {
			return t, false
		}
		r := slice[i]
		i--
		return r, true
	}
}

type KV[K comparable, V any] struct {
	Key K
	Val V
}

func Map[K comparable, V any](m map[K]V) Iter[KV[K, V]] {
	keys := make([]KV[K, V], 0, len(m))
	for k, v := range m {
		keys = append(keys, KV[K, V]{
			Key: k,
			Val: v,
		})
	}

	return Slice(keys)
}
