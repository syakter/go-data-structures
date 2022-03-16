package trie

type Trie[V any] struct {
	n    int
	root *node[V]
}

type node[V any] struct {
	c                byte
	left, mid, right *node[V]
	val              V
	valid            bool
}

func New[V any]() *Trie[V] {
	return &Trie[V]{}
}

func (t *Trie[V]) Contains(key string) bool {
	if len(key) == 0 {
		return false
	}
	_, ok := t.Get(key)
	return ok
}

func (t *Trie[V]) Get(key string) (v V, ok bool) {
	if len(key) == 0 {
		return v, false
	}
	x := t.get(t.root, key, 0)
	if x == nil || !x.valid {
		return v, false
	}
	return x.val, true
}

func (t *Trie[V]) get(x *node[V], key string, d int) *node[V] {
	if x == nil || len(key) == 0 {
		return nil
	}
	c := key[d]
	if c < x.c {
		return t.get(x.left, key, d)
	} else if c > x.c {
		return t.get(x.right, key, d)
	} else if d < len(key)-1 {
		return t.get(x.mid, key, d+1)
	} else {
		return x
	}
}

func (t *Trie[V]) Remove(key string) {
	if len(key) == 0 || !t.Contains(key) {
		return
	}
	var v V
	t.n--
	t.root = t.Put(t.root, key, v, 0, false)
}

func (t *Trie[V]) Put(x *node[V], key string, val V, d int, valid bool) *node[V] {
	c := key[d]
	if x == nil {
		x = &node[V]{
			c: c,
		}
	}
	if c < x.c {
		x.left = t.Put(x.left, key, val, d, valid)
	} else if c > x.c {
		x.right = t.Put(x.right, key, val, d, valid)
	} else if d < len(key)-1 {
		x.mid = t.Put(x.mid, key, val, d+1, valid)
	} else {
		x.val = val
		x.valid = valid
	}
	return x
}

func (t *Trie[V]) LongestPrefix(query string) string {
	if len(query) == 0 {
		return ""
	}
	length := 0
	x := t.root
	i := 0
	for x != nil && i < len(query) {
		c := query[i]
		if c < x.c {
			x = x.left
		} else if c > x.c {
			x = x.right
		} else {
			i++
			if x.valid {
				length = i
			}
			x = x.mid
		}
	}
	return query[:length]
}

func (t *Trie[V]) Keys() (queue []string) {
	return t.collect(t.root, "", queue)
}

func (t *Trie[V]) KeysWithPrefix(prefix string) (queue []string) {
	if len(prefix) == 0 {
		return t.Keys()
	}
	x := t.get(t.root, prefix, 0)
	if x == nil {
		return nil
	}
	if x.valid {
		queue = []string{prefix}
	}
	return t.collect(x.mid, prefix, queue)
}

func (t *Trie[V]) collect(x *node[V], prefix string, queue []string) []string {
	if x == nil {
		return queue
	}
	queue = t.collect(x.left, prefix, queue)
	if x.valid {
		queue = append(queue, prefix+string(x.c))
	}
	queue = t.collect(x.mid, prefix+string(x.c), queue)
	if len(prefix) > 0 {
		prefix = prefix[:len(prefix)-1]
	}
	return t.collect(x.right, prefix, queue)
}
