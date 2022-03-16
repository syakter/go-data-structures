package hashmap_test

import (
	"fmt"
	"math/rand"
	"testing"

	g "github.com/syakter/go-data-structures/generic"
	"github.com/syakter/go-data-structures/hashmap"
)

func checkeq[K any, V comparable](cm *hashmap.Map[K, V], get func(k K) (V, bool), t *testing.T) {
	cm.Iter().For(func(kv hashmap.KV[K, V]) {
		if ov, ok := get(kv.Key); !ok {
			t.Fatalf("key %v should exist", kv.Key)
		} else if kv.Val != ov {
			t.Fatalf("value mismatch: %v != %v", kv.Val, ov)
		}
	})
}

func TestCrossCheck(t *testing.T) {
	stdm := make(map[uint64]uint32)
	cowm := hashmap.NewMap[uint64, uint32](1, g.Equals[uint64], g.HashUint64)

	const nops = 1000

	for i := 0; i < nops; i++ {
		key := uint64(rand.Intn(100))
		val := rand.Uint32()
		op := rand.Intn(2)

		switch op {
		case 0:
			stdm[key] = val
			cowm.Put(key, val)
		case 1:
			var del uint64
			for k := range stdm {
				del = k
				break
			}
			delete(stdm, del)
			cowm.Remove(del)
		}

		checkeq(cowm, func(k uint64) (uint32, bool) {
			v, ok := stdm[k]
			return v, ok
		}, t)
	}
}

func TestCopy(t *testing.T) {
	orig := hashmap.NewMap[uint64, uint32](1, g.Equals[uint64], g.HashUint64)

	for i := uint32(0); i < 10; i++ {
		orig.Put(uint64(i), i)
	}

	cpy := orig.Copy()

	checkeq(cpy, orig.Get, t)

	cpy.Put(0, 42)

	if v, _ := cpy.Get(0); v != 42 {
		t.Fatal("didn't get 42")
	}
}

func Example() {
	m := hashmap.NewMap[string, int](1, g.Equals[string], g.HashString)
	m.Put("foo", 42)
	m.Put("bar", 13)

	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("baz"))

	m.Remove("foo")

	fmt.Println(m.Get("foo"))
}
