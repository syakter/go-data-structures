package list_test

import (
	"fmt"

	"github.com/syakter/go-data-structures/list"
)

func Example() {
	l := list.New[int]()
	l.PushBack(0)
	l.PushBack(1)
	l.PushBack(2)

	l.Front.Iter().For(func(i int) {
		fmt.Println(i)
	})
}
